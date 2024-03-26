package router

import (
	"github.com/gin-gonic/gin"
	"todopoint/member/router/controller"
)

func NewMemberRouter(controller *controller.MemberController) *gin.Engine {
	engine := gin.Default()
	router := engine.Group("/members")
	{
		router.POST("/register", controller.RegisterMember)
		router.GET("/:userId/valid", controller.ValidateMember)
	}
	return engine
}
