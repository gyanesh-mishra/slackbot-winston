package health

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetResponse defines the HTTP response model for an incoming GET request
type GetResponse struct {
	Success bool `json:"success,omitempty"`
}

// HandleGet handles the incoming HTTP GET request
func HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var response GetResponse
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
