package slackevents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	questionAnswerDAO "github.com/gyanesh-mishra/slackbot-winston/internal/dao/questionanswer"
	"github.com/gyanesh-mishra/slackbot-winston/internal/helpers"

	"github.com/gyanesh-mishra/slackbot-winston/config"

	"github.com/julienschmidt/httprouter"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// HandlePost handles the incoming HTTP POST request
func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var configuration = config.GetConfig()
	var api = configuration.Slack.Client

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	// Check if the request is valid and coming from Events API
	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: configuration.Slack.VerificationToken}))
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Slack URL verification method
	if eventsAPIEvent.Type == slackevents.URLVerification {
		handleURLVerificationEvent(w, body)
	}

	// Slack messages in channel
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		handleCallbackEvent(w, api, eventsAPIEvent.InnerEvent)
	}
}

// handleUrlVerificationEvent handles the bot verification request when a slack integration happens
func handleURLVerificationEvent(w http.ResponseWriter, body string) {
	var r *slackevents.ChallengeResponse
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text")
	w.Write([]byte(r.Challenge))
}

// handleCallbackEvent handles the Slack messages event
func handleCallbackEvent(w http.ResponseWriter, api *slack.Client, innerEvent slackevents.EventsAPIInnerEvent) {

	switch ev := innerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:

		// Remove user mentions
		question := helpers.RemoveUserMention(ev.Text)

		// Sanitize input
		question = helpers.ExtractQuestionFromMessage(question)

		// Get the answer from the database
		answer, err := questionAnswerDAO.GetAnswerByQuestion(question)

		// If no answers were found, prompt channel to help
		if err != nil {
			question = fmt.Sprintf("%s", question)
			responseAttachment := slack.MsgOptionAttachments(getAnswerNotFoundAttachment(question))
			response := fmt.Sprintf("I can't seem to remember atm, :thinking_face: Can someone help me out?")
			api.PostMessage(ev.Channel, slack.MsgOptionText(response, false), responseAttachment)
			return

		}

		// Respond with answer found in the database
		response := "Found an answer! :smile:"
		attachmentMessage := fmt.Sprintf("%s \n _%s_", question, answer)
		responseAttachment := slack.MsgOptionAttachments(getAnswerFoundAttachment(attachmentMessage))
		api.PostMessage(ev.Channel, slack.MsgOptionText(response, false), responseAttachment)
		return
	default:
		log.Print(fmt.Sprintf("Uncaught Event: %+v\n", ev))
	}
}
