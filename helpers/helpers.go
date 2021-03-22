package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/nehal1992/Go-Clean-Architecture/serviceErrors"
)

func JsonResponse(w http.ResponseWriter, response interface{}, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
	return
}
func JsonErrorResponse(w http.ResponseWriter, err error, code int) {
	JsonResponse(w, serviceErrors.Error{err.Error()}, code)

}
