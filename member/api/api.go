package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"todopoint/member/internal/pkg/model"
)

var members map[int]model.Member

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/members", GetMemberListHandler).Methods("GET")
	mux.HandleFunc("/members", PostMemberHandler).Methods("POST")
	// test data
	members = make(map[int]model.Member)
	members[0] = model.Member{
		Id:       1,
		UserId:   "helloworld",
		UserName: "hello",
		UserPw:   "hellopw",
	}
	members[1] = model.Member{Id: 2, UserId: "helloworld2", UserName: "hello2", UserPw: "hellopw2"}

	return mux
}

func GetMemberListHandler(w http.ResponseWriter, r *http.Request) {
	type Members []model.Member // list
	list := make(Members, 0)
	for _, member := range members {
		list = append(list, member)
	}
	writeJson(w, http.StatusOK, list)

}

func PostMemberHandler(w http.ResponseWriter, r *http.Request) {
	var member model.Member
	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
