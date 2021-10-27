package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohammadraasel/go-api/internal/comment"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetAllComments(wr http.ResponseWriter, r *http.Request) {

	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(wr, "Error retrieving comments")
	}

	fmt.Fprintf(wr, "%+v", comments)
}

func (h *Handler) CreateComment(wr http.ResponseWriter, r *http.Request) {

	comment, err := h.Service.CreateComment(comment.Comment{
		Slug: "/",
	})
	if err != nil {
		fmt.Fprintf(wr, "Failed to create comment")
	}

	fmt.Fprintf(wr, "%+v", comment)
}

func (h *Handler) GetComment(wr http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		fmt.Fprintf(wr, "Unable to parse uint from id")
	}

	comment, err := h.Service.GetComment(uint(id))
	if err != nil {
		fmt.Fprintf(wr, "Error retrieving comment by id")
	}

	fmt.Fprintf(wr, "%+v", comment)

}

func (h *Handler) HealthCheck(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "I am alive!")
}

func (h *Handler) SetupRoutes() {
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comments", h.CreateComment).Methods("POST")
	h.Router.HandleFunc("/api/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/health", h.HealthCheck).Methods("GET")
}
