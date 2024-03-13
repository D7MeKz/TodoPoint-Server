package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func GetAllMembersController(w http.ResponseWriter, r *http.Request) {
	members, err := service.NewMemberRepo(r.Context()).GetAllMembers()
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}
	utils.Return(w, true, http.StatusOK, nil, members)
}

func GetMemberByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}
	member, err := service.NewMemberRepo(r.Context()).GetMemberByID(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}
	utils.Return(w, true, http.StatusOK, nil, member)

}

func DeleteMemberController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewMemberRepo(r.Context()).DeleteMember(id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}
	utils.Return(w, true, http.StatusOK, nil, deletedID)
}

func UpdateMemberController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	var targetMember model.Member
	err = json.NewDecoder(r.Body).Decode(&targetMember)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	modifiedMember, err := service.NewMemberRepo(r.Context()).UpdateMember(id, targetMember)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}
	utils.Return(w, true, http.StatusOK, nil, modifiedMember)
}
