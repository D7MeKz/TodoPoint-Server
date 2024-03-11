package controller

import (
	"encoding/json"
	"net/http"
	"todopoint/member/internal/model"
	"todopoint/member/internal/service"
	"todopoint/member/internal/utils"
)

func CreateMemberController(w http.ResponseWriter, r *http.Request) {
	var newMember model.Member
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	createdMember, err := service.NewMemberRepo(r.Context()).CreateMember(newMember)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}
	utils.Return(w, true, http.StatusOK, nil, createdMember)
}
