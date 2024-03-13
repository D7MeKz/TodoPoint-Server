package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"todopoint/member/internal/controller"
)

func registerMemberRouter(r *mux.Router) {
	memberRouter := r.PathPrefix("/members").Subrouter()
	memberRouter.HandleFunc("/register", controller.CreateMemberController).Methods(http.MethodPost)
	memberRouter.HandleFunc("/{id}", controller.GetMemberByIDController).Methods(http.MethodGet)
	memberRouter.HandleFunc("/all", controller.GetAllMembersController).Methods(http.MethodGet)
	memberRouter.HandleFunc("/delete/{id}", controller.DeleteMemberController).Methods(http.MethodGet)
	memberRouter.HandleFunc("/update/{id}", controller.UpdateMemberController).Methods(http.MethodPost)
}
