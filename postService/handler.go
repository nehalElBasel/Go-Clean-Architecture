package postService

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nehal1992/Go-Clean-Architecture/helpers"
	"github.com/nehal1992/Go-Clean-Architecture/models"
)

var getPostId = func(r *http.Request) (int, error) {
	post_id_url := mux.Vars(r)["id"]
	return strconv.Atoi(post_id_url)

}

func (p *PostService) ListHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := p.List(r.Context())
	if err != nil {
		helpers.JsonErrorResponse(w, err, http.StatusInternalServerError)
	}

	helpers.JsonResponse(w, posts, http.StatusAccepted)

}

func (p *PostService) GetHandler(w http.ResponseWriter, r *http.Request) {
	post_id, err := getPostId(r)
	if err != nil {
		helpers.JsonErrorResponse(w, err, http.StatusInternalServerError)
	}
	post, err := p.Get(r.Context(), post_id)
	if err != nil {
		helpers.JsonErrorResponse(w, err, http.StatusInternalServerError)
	}
	helpers.JsonResponse(w, post, http.StatusAccepted)
}

func (p *PostService) AddHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	err := p.Add(r.Context(), post)
	if err != nil {
		helpers.JsonErrorResponse(w, err, http.StatusInternalServerError)
	}
	helpers.JsonResponse(w, "added successfully", http.StatusAccepted)
}

func (p *PostService) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	post_id, err := getPostId(r)
	if err != nil {
		helpers.JsonErrorResponse(w, err, http.StatusInternalServerError)
	}
	json.NewDecoder(r.Body).Decode(&post)
	err = p.Update(r.Context(), post_id, post)
	if err != nil {
		helpers.JsonErrorResponse(w, err, http.StatusInternalServerError)
	}
	helpers.JsonResponse(w, "updated successfully", http.StatusAccepted)
}

func (p *PostService) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	post_id, err := getPostId(r)
	if err != nil {
		helpers.JsonErrorResponse(w, err, http.StatusInternalServerError)
	}
	res, err := p.Delete(r.Context(), post_id)
	if err != nil {
		helpers.JsonErrorResponse(w, err, http.StatusInternalServerError)
	}
	helpers.JsonResponse(w, res, http.StatusAccepted)
}
