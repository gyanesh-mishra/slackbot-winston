package slackevents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gyanesh-mishra/slackbot-winston/internal/constants"
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
	// TODO: Update tokens auth to OAuth Flow
	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body),
		slackevents.OptionVerifyToken(
			&slackevents.TokenComparator{VerificationToken: configuration.Slack.VerificationToken},
		),
	)
	if e != nil {
		log.Printf("Error parsing slack event %+v\n", e)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Slack URL verification method
	if eventsAPIEvent.Type == slackevents.URLVerification {
		handleURLVerificationEvent(w, body)
	}

	// Slack messages in channel
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		handleCallbackEvent(w, api, eventsAPIEvent)
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
func handleCallbackEvent(w http.ResponseWriter, api *slack.Client, event slackevents.EventsAPIEvent) {

	switch ev := event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:

		// Remove user mentions
		question := helpers.RemoveUserMention(ev.Text)

		// Sanitize input
		question = helpers.ExtractQuestionFromMessage(question)
		// Exit if there is no question
		if question == "" {
			response := helpers.GetRandomStringFromSlice(constants.GreetingMessages)
			api.PostMessage(ev.Channel, slack.MsgOptionText(response, false))
			return
		}

		// Any question that has less than 10 characters is probably jumbled
		// TODO: Change this to count words or do better
		if len(question) < 10 {
			response := "I don't know how to answer that :thinking_face:"
			api.PostMessage(ev.Channel, slack.MsgOptionText(response, false))
			return
		}

		// Get the answer from the database
		result, err := questionAnswerDAO.GetByQuestion(question, event.TeamID)

		// If no answers were found, prompt channel to help
		if err != nil {
			question = fmt.Sprintf("%s", question)
			responseAttachment := slack.MsgOptionAttachments(getAnswerNotFoundAttachment(question))
			response := helpers.GetRandomStringFromSlice(constants.AnswerNotFoundMessages)
			api.PostMessage(ev.Channel, slack.MsgOptionText(response, false), responseAttachment)
			return

		}

		// Respond with answer found in the database
		response := helpers.GetRandomStringFromSlice(constants.AnswerFoundMessages)
		resultLastUpdated := time.Now().UTC().Sub(result.LastUpdated).Round(time.Hour)
		attachmentMessage := fmt.Sprintf("%s \n _%s_ \n \n Last updated by %s %s ago", question, result.Answer, result.LastUpdatedBy, resultLastUpdated)
		responseAttachment := slack.MsgOptionAttachments(getAnswerFoundAttachment(attachmentMessage))
		api.PostMessage(ev.Channel, slack.MsgOptionText(response, false), responseAttachment)
		return
	case *slackevents.MessageEvent:
		if ev.SubType == "bot_add" {
			fmt.Printf("BOT ADDED : %+v\n", event)
		}
	default:
		log.Print(fmt.Sprintf("Uncaught Event: %+v\n", ev))
	}
}
