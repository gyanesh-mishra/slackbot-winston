package constants

// Attachment constants
const (
	AnswerNotFoundAttachmentID    = "answer_not_found_attachment_id"
	AnswerNotFoundVolunteerAction = "answer_not_found_volunteer_action"
	AnswerUserInputDialogID       = "answer_user_input_dialog_id"

	AnswerUserInputDialogInput = "answer_user_input_dialog_input"

	AnswerFoundUpdateAttachmentID = "answer_found_update_attachment_id"
	AnswerFoundUpdateAction       = "answer_found_update_action"
)

// GreetingMessages contains generic greeting messages.
var GreetingMessages = []string{
	"Hello there! :wave:",
	"Did someone miss me? :star-struck:",
	"Ask! and Ye shall receive the Truth!!!\n_Unless ofcourse I don't know what you're talking about_ :sweat_smile:",
	"I'm here :raising_hand:",
}

// AnswerNotFoundMessages contains messages for when an answer is not in the DB.
var AnswerNotFoundMessages = []string{
	"I can't seem to remember ath the moment, :thinking_face: Can someone help me out?",
	"I'm stumped, can someone jog my memory? :sweat:",
	"Aw man, I knew I should've taken it easy last night :stuck_out_tongue:, Can anyone help?",
	"I don't know _yet_, but I will soon!",
}

// AnswerFoundMessages contains messages for when an answer is in the DB.
var AnswerFoundMessages = []string{
	"Found an answer! :smile:",
	"Here you go! :smile:",
	"Is this correct? :raising_hand:",
	"This is what I know, is it still accurate?",
}

// NewAnswerMessages contains messages for when a new answer is added to the DB.
var NewAnswerMessages = []string{
	"I just learnt something new!",
	"Is this considered machine learning? :wink:",
	"I feel smarter now!",
}
