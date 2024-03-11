package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"todopoint/member/internal/controller"
)

func registerMemberRouter(r *mux.Router) {
	memberRouter := r.PathPrefix("/member").Subrouter()
	memberRouter.HandleFunc("/", controller.CreateMemberController).Methods(http.MethodPost)
}
