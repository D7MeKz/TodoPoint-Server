package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"todopoint/member/types"
)

var members map[int]types.Member

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func MakeWebHanlder() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/members", GetMemberListHandler).Methods("GET")
	// test data
	members = make(map[int]types.Member)
	members[0] = types.Member{1, "helloworld", "hello", "hellopw"}
	members[1] = types.Member{2, "helloworld2", "hello2", "hellopw2"}

	return mux
}

func GetMemberListHandler(w http.ResponseWriter, r *http.Request) {
	type Members []types.Member // list
	list := make(Members, 0)
	for _, member := range members {
		list = append(list, member)
	}
	writeJson(w, http.StatusOK, list)

}
