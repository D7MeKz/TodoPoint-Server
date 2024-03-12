package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"todopoint/member/internal/controller"
)

func registerMemberRouter(r *mux.Router) {
	memberRouter := r.PathPrefix("/member").Subrouter()
	memberRouter.HandleFunc("/", controller.CreateMemberController).Methods(http.MethodPost)
	memberRouter.HandleFunc("/{id}", controller.GetMemberByIDController).Methods(http.MethodGet)
	memberRouter.HandleFunc("/", controller.GetAllMembersController).Methods(http.MethodGet)
	memberRouter.HandleFunc("/{id}", controller.DeleteMemberController).Methods(http.MethodDelete)
	memberRouter.HandleFunc("/{id}", controller.UpdateMemberController).Methods(http.MethodPut)
}
