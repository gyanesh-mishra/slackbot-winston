package slackevents

import (
	"github.com/gyanesh-mishra/slackbot-winston/internal/constants"
	"github.com/nlopes/slack"
)

func getAnswerNotFoundAttachment(question string) slack.Attachment {
	return slack.Attachment{
		Text:       question,
		Color:      "#f9a41b",
		CallbackID: constants.AnswerNotFoundAttachmentID,
		Actions: []slack.AttachmentAction{
			{
				Name:  constants.AnswerNotFoundVolunteerAction,
				Text:  "I can :raising_hand:",
				Type:  "button",
				Style: "primary",
				Value: constants.AnswerNotFoundVolunteerAction,
			},
		},
	}
}

func getAnswerFoundAttachment(message string) slack.Attachment {
	return slack.Attachment{
		Text:       message,
		Color:      "#1818c4",
		CallbackID: constants.AnswerFoundUpdateAttachmentID,
		Actions: []slack.AttachmentAction{
			{
				Name:  constants.AnswerFoundUpdateAction,
				Text:  "Update answer :pencil:",
				Type:  "button",
				Style: "primary",
				Value: constants.AnswerFoundUpdateAction,
			},
		},
	}
}
