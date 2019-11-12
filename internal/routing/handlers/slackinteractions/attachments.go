package slackinteractions

import (
	"github.com/nlopes/slack"
)

func getUserAnswerForQuestionDialog(question string, triggerID string, callbackID string) slack.Dialog {
	return slack.Dialog{
		TriggerID:   triggerID,
		CallbackID:  callbackID,
		Title:       "Submit your answer",
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
