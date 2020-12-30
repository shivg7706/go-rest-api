package controller

import (
	"net/http"

	"github.com/shivg7706/rest-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To dhinchak API")

}
