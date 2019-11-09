package slackinteractions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gyanesh-mishra/slackbot-winston/config"
	"github.com/gyanesh-mishra/slackbot-winston/internal/constants"
	questionAnswerDAO "github.com/gyanesh-mishra/slackbot-winston/internal/dao/questionanswer"
	"github.com/julienschmidt/httprouter"
	"github.com/nlopes/slack"
)

// HandlePost handles the incoming HTTP POST request
func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var configuration = config.GetConfig()
	var slackAPI = configuration.Slack.Client

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

	log.Print(fmt.Sprintf("Interaction Endpoint log: %+v\n", payload))

	switch payload.CallbackID {
	case constants.AnswerNotFoundAttachmentID:
		// Answer was not found in the database, get answer from a volunteer

		// Update the original message and append helper name as attachment
		responseMessage := payload.OriginalMessage
		responseMessage.ResponseType = "in_channel"
		responseMessage.ReplaceOriginal = true
		responseMessage.Attachments[0].Actions = nil
		responseMessage.Attachments = append(responseMessage.Attachments, slack.Attachment{
			Text:       fmt.Sprintf("<@%s> is helping me! :fire: ", payload.User.ID),
			Color:      "#062F67",
			CallbackID: "untracked_event",
		})

		// Return new response to be updated
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseMessage)

		// Prompt the volunteer for answer
		dialog := getUserAnswerForQuestionDialog(payload.OriginalMessage.Attachments[0].Text, payload.TriggerID, constants.AnswerUserInputDialogID)
		defer slackAPI.OpenDialog(payload.TriggerID, dialog)
		return
	case constants.AnswerUserInputDialogID:
		// User has provided answer for the question, store it in DB
		for question, answer := range payload.DialogSubmissionCallback.Submission {
			_, err := questionAnswerDAO.Add(question, answer)
			if err != nil {
				fmt.Printf("Error adding QnA : %+v\n", err)
			}
			// Post on the channel about learning new information
			defer sendUserConfirmation(question, answer, payload.Channel.ID, slackAPI)
		}
	case constants.AnswerFoundUpdateAttachmentID:

		// Update the original message and append helper name as attachment
		responseMessage := payload.OriginalMessage
		responseMessage.ResponseType = "in_channel"
		responseMessage.ReplaceOriginal = true
		responseMessage.Attachments[0].Actions = nil
		responseMessage.Attachments = append(responseMessage.Attachments, slack.Attachment{
			Text:       fmt.Sprintf("<@%s> is updating this answer :raised_hands: ", payload.User.ID),
			CallbackID: "untracked_event",
		})

		// Return new response to be updated
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseMessage)

		// Prompt the volunteer for answer
		dialog := getUserAnswerForQuestionDialog(payload.OriginalMessage.Attachments[0].Text, payload.TriggerID, constants.AnswerUserInputDialogID)
		defer slackAPI.OpenDialog(payload.TriggerID, dialog)
		return
	default:
		log.Print(fmt.Sprintf("Unhandled Callback ID: %+v\n", payload))
	}

}

func sendUserConfirmation(question string, answer string, channel string, slackAPI *slack.Client) {
	responseAttachment := slack.MsgOptionAttachments(slack.Attachment{
		Text:       fmt.Sprintf("Question : %s \n Answer : %s", question, answer),
		Color:      "#062F67",
		CallbackID: "untracked_event",
	})
	responseTitle := "I just learnt something new!"
	slackAPI.PostMessage(channel, slack.MsgOptionText(responseTitle, false), responseAttachment)
}
