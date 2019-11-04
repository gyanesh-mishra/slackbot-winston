package slackevents

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"

	questionAnswerDAO "github.com/gyanesh-mishra/slackbot-winston/internal/dao/questionanswer"

	"github.com/gyanesh-mishra/slackbot-winston/config"

	"github.com/julienschmidt/httprouter"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// GetResponse defines the HTTP response model for an incoming GET request
type GetResponse struct {
	message string
}

// HandleGet handles the incoming HTTP GET request
func HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var response GetResponse
	response.message = "Hello!"
	json.NewEncoder(w).Encode(response)
}

// HandlePost handles the incoming HTTP POST request
func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var configuration = config.GetConfig()

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
		var api = slack.New(configuration.Slack.BotToken)
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

		// Remove any user mentions
		re := regexp.MustCompile(`<@[^>]*>`)
		message := re.ReplaceAllString(ev.Text, "")

		response, err := questionAnswerDAO.GetAnswerByQuestion(message)
		if err != nil {
			response = "I don't understand, please educate me"
		}
		api.PostMessage(ev.Channel, slack.MsgOptionText(response, false))
	default:
		//fmt.Printf("Uncaught Event %+v", ev)
	}
}
