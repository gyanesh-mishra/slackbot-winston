package slackinteractions

import (
	"github.com/gyanesh-mishra/slackbot-winston/internal/constants"
	"github.com/nlopes/slack"
)

func getUserAnswerForQuestionDialog(question string, triggerID string) slack.Dialog {
	return slack.Dialog{
		TriggerID:   triggerID,
		CallbackID:  constants.AnswerUserInputDialogID,
		Title:       "I owe ya one!",
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
