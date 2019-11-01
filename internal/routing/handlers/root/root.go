package root

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gyanesh-mishra/slackbot-winston/config"
	"gopkg.in/jdkato/prose.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// HandleGet exported
func HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var response GetResponse
	response.message = "Hello!"
	json.NewEncoder(w).Encode(response)
}

// HandlePost exported
func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var configuration = config.GetConfig()
	var api = slack.New(configuration.Slack.BotToken)

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: configuration.Slack.VerificationToken}))
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent

		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			// Remove any user mentions
			re := regexp.MustCompile(`<@[^>]*>`)
			message := re.ReplaceAllString(ev.Text, "")

			// Parse input and break into keywords
			keywords := []string{}
			doc, _ := prose.NewDocument(message, prose.WithExtraction(false))
			// Iterate over the doc's tokens:
			for _, tok := range doc.Tokens() {
				fmt.Println(tok.Text, tok.Tag)
				if tok.Tag == "NN" || tok.Tag == "VBG" || tok.Tag == "NNS" {
					keywords = append(keywords, tok.Text)
				}
			}
			api.PostMessage(ev.Channel, slack.MsgOptionText(message, false))
		default:
			//fmt.Println(ev)
		}
	}
}
