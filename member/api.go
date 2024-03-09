package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func MakeWebHanlder() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/members", GetMemberListHandler).Methods("GET")
	// test data
	members = make(map[int]Member)
	members[0] = Member{1, "helloworld", "hello", "hellopw"}
	members[1] = Member{2, "helloworld2", "hello2", "hellopw2"}
	lastId = 2

	return mux
}

func GetMemberListHandler(w http.ResponseWriter, r *http.Request) {
	type Members []Member // list
	list := make(Members, 0)
	for _, member := range members {
		list = append(list, member)
	}
	writeJson(w, http.StatusOK, list)

}
