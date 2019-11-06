package slackevents

import (
	"github.com/gyanesh-mishra/slackbot-winston/internal/constants"
	"github.com/nlopes/slack"
)

var helpAnswerAttachment = slack.Attachment{
	Text:       "Can someone help me?",
	Color:      "#f9a41b",
	CallbackID: constants.HelpAnswerCallBackID,
	Actions: []slack.AttachmentAction{
		{
			Name:  constants.HelpAnswerYes,
			Text:  "Yes",
			Type:  "button",
			Style: "primary",
			Value: constants.HelpAnswerYes,
		},
		{
			Name:  constants.HelpAnswerNo,
			Text:  "No",
			Type:  "button",
			Style: "danger",
			Value: constants.HelpAnswerNo,
		},
	},
}
