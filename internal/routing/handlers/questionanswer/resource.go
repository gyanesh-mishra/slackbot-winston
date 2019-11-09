package questionanswer

import (
	"bytes"
	"encoding/json"
	"net/http"

	questionAnswerDAO "github.com/gyanesh-mishra/slackbot-winston/internal/dao/questionanswer"
	"github.com/gyanesh-mishra/slackbot-winston/internal/helpers"

	"github.com/julienschmidt/httprouter"
)

// HandleGet handles incoming HTTP GET requests
func HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := questionAnswerDAO.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(res)
}

// PostRequestQuestionAnswer defines the body of an incoming HTTP POST request
type PostRequestQuestionAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// HandlePost handles the incoming HTTP POST request
func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get Request body
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	// Decode request body into the PostRequestQuestionAnswer struct
	var data PostRequestQuestionAnswer
	err := json.Unmarshal([]byte(buf.String()), &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Sanitize the question before inserting
	question := helpers.ExtractQuestionFromMessage(data.Question)
	res, err := questionAnswerDAO.Add(question, data.Answer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(res)

}
