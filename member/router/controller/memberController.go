package controller

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/errorutils"
	"todopoint/common/webutils"
	"todopoint/member/data/request"
	"todopoint/member/service"
)

type MemberController struct {
	service service.MemberService
}

func NewMemberController(s service.MemberService) *MemberController {
	return &MemberController{
		service: s,
	}
}

func (controller *MemberController) RegisterMember(ctx *gin.Context, req request.RegisterReq) {
	req = request.RegisterReq{}
	err := ctx.ShouldBindJSON(req)

	if err != nil {
		box := errorutils.NewErrorBox(errorutils.INVALID_JSON_FORMAT, err, "")
		errorutils.ErrorResponse(ctx, box)
		return
	}

	// Create member
	errBox := controller.service.CreateMember(ctx, req)
	if errBox != nil {
		errorutils.ErrorResponse(ctx, errBox)
		return
	}

	webutils.Success(ctx)
}

//func CreateMemberController(w http.ResponseWriter, r *http.Request) {
//	var newMember model.Member
//	err := json.NewDecoder(r.Body).Decode(&newMember)
//	if err != nil {
//		utils.Return(w, false, http.StatusBadRequest, err, nil)
//		return
//	}
//
//	createdMember, err := service.NewMemberRepo(r.Context()).CreateMember(newMember)
//	if err != nil {
//		utils.Return(w, false, http.StatusInternalServerError, err, nil)
//	}
//	utils.Return(w, true, http.StatusOK, nil, createdMember)
//}
//
//func GetAllMembersController(w http.ResponseWriter, r *http.Request) {
//	members, err := service.NewMemberRepo(r.Context()).GetAllMembers()
//	if err != nil {
//		utils.Return(w, false, http.StatusInternalServerError, err, nil)
//	}
//	utils.Return(w, true, http.StatusOK, nil, members)
//}
//
//func GetMemberByIDController(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		utils.Return(w, false, http.StatusBadRequest, err, nil)
//		return
//	}
//	member, err := service.NewMemberRepo(r.Context()).GetMemberByID(id)
//	if err != nil {
//		utils.Return(w, false, http.StatusInternalServerError, err, nil)
//		return
//	}
//	utils.Return(w, true, http.StatusOK, nil, member)
//
//}
//
//func DeleteMemberController(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		utils.Return(w, false, http.StatusBadRequest, err, nil)
//		return
//	}
//
//	deletedID, err := service.NewMemberRepo(r.Context()).DeleteMember(id)
//	if err != nil {
//		utils.Return(w, false, http.StatusInternalServerError, err, nil)
//		return
//	}
//	utils.Return(w, true, http.StatusOK, nil, deletedID)
//}
//
//func UpdateMemberController(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		utils.Return(w, false, http.StatusBadRequest, err, nil)
//		return
//	}
//
//	var targetMember model.Member
//	err = json.NewDecoder(r.Body).Decode(&targetMember)
//	if err != nil {
//		utils.Return(w, false, http.StatusBadRequest, err, nil)
//		return
//	}
//
//	modifiedMember, err := service.NewMemberRepo(r.Context()).UpdateMember(id, targetMember)
//	if err != nil {
//		utils.Return(w, false, http.StatusInternalServerError, err, nil)
//		return
//	}
//	utils.Return(w, true, http.StatusOK, nil, modifiedMember)
//}
//
//func ValidationController(w http.ResponseWriter, r *http.Request) {
//	log.Println("Start Validate Member validation")
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		utils.Return(w, false, http.StatusBadRequest, err, nil)
//		return
//	}
//	_, err = service.NewMemberRepo(r.Context()).GetMemberByID(id)
//	if err != nil {
//
//		utils.Return(w, false, http.StatusUnauthorized, err, nil)
//		return
//	}
//
//	utils.Return(w, true, http.StatusOK, nil, nil)
//}
