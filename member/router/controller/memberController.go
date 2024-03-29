package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	wu "todopoint/common/webutils"
	"todopoint/member/data/request"
	"todopoint/member/data/response"
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

func (controller *MemberController) RegisterMember(ctx *gin.Context) {
	req := request.RegisterReq{}
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		wu.ErrorFunc(ctx, wu.NewError(wu.INVALID_JSON_FORMAT, err))
		return
	}

	// Create member
	mem := controller.service.CreateMember(ctx, req)
	mid := response.MemberId{MemberId: mem.ID}
	wu.SuccessWith(ctx, &mid)
	return
}

func (controller *MemberController) LoginMember(ctx *gin.Context) {
	req := request.LoginReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		wu.ErrorFunc(ctx, wu.NewError(wu.INVALID_JSON_FORMAT, err))
		return
	}
	// login Member
	memId, err2 := controller.service.LoginMember(ctx, req)
	fmt.Print(memId)
	if err != nil || memId == -1 {
		wu.ErrorFunc(ctx, err2)
		return
	}
	res := response.MemberId{MemberId: memId}
	wu.SuccessWith(ctx, res)

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
