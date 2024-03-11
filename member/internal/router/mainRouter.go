package router

import "github.com/gorilla/mux"

func RegisterMainRouter(r *mux.Router) {
	registerMemberRouter(r)
}
