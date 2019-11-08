package slackinteractions

import (
	"github.com/gyanesh-mishra/slackbot-winston/internal/constants"
	"github.com/nlopes/slack"
)

func getUserAnswerForQuestionDialog(question string, triggerID string) slack.Dialog {
	return slack.Dialog{
		TriggerID:   triggerID,
		CallbackID:  constants.AnswerUserInputDialogID,
		Title:       "Thanks for sharing!",
		SubmitLabel: "Submit",
		Elements: []slack.DialogElement{
			slack.DialogInput{
				Type:        "text",
				Label:       question,
				Name:        question,
				Placeholder: "Type your answer here",
				Optional:    false,
			},
		},
	}
}
