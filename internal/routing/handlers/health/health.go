package health

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetResponse model exported
type GetResponse struct {
	Success bool `json:"success,omitempty"`
}

// HandleGet exported
func HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var response GetResponse
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
