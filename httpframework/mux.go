package httpframework

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxHttp struct{}

func NewMuxHttp() HttpFramework {
	return &muxHttp{}
}
func (muxhttp *muxHttp) HandelRoutes(m *Modules) {
	r := mux.NewRouter()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome test")
	})
	postRoutes := r.PathPrefix("/posts").Subrouter().StrictSlash(true)
	postRoutes.HandleFunc("/", m.PostService.ListHandler).Methods("Get")
	postRoutes.HandleFunc("/", m.PostService.AddHandler).Methods("Post")
	postRoutes.HandleFunc("/{id}", m.PostService.GetHandler).Methods("Get")
	postRoutes.HandleFunc("/{id}", m.PostService.UpdateHandler).Methods("Put")
	postRoutes.HandleFunc("/{id}", m.PostService.DeleteHandler).Methods("Delete")
	http.ListenAndServe(":8080", r)
}
