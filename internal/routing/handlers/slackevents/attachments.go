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
				Name:  constants.AnswerNotFoundVolunteer,
				Text:  "I can :raising_hand:",
				Type:  "button",
				Style: "primary",
				Value: constants.AnswerNotFoundVolunteer,
			},
		},
	}

}
