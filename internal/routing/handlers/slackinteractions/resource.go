package slackinteractions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gyanesh-mishra/slackbot-winston/config"
	"github.com/gyanesh-mishra/slackbot-winston/internal/constants"
	"github.com/julienschmidt/httprouter"
	"github.com/nlopes/slack"
)

// HandlePost handles the incoming HTTP POST request
func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var configuration = config.GetConfig()

	var payload slack.InteractionCallback
	err := json.Unmarshal([]byte(r.FormValue("payload")), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if payload.Token != configuration.Slack.VerificationToken {
		w.WriteHeader(http.StatusUnauthorized)
	}

	switch payload.CallbackID {
	case constants.HelpAnswerCallBackID:
		response := fmt.Sprintf("<@%s> is helping me! :fire: ", payload.User.ID)
		w.Write([]byte(response))
	default:
		log.Print(fmt.Sprintf("Unhandled Callback ID: %+v\n", payload))
	}

}
