package general

import (
	"football_api/helpers"
	"net/http"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJson(w, 200, "OK")
	return
}
