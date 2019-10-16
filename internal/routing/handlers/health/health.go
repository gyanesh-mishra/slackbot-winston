package root

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Success!\n")
}
