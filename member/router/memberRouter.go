package router

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/netutils/middleware"
	"todopoint/member/router/controller"
)

func NewMemberRouter(controller *controller.MemberController) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.ErrorHandler())

	router := engine.Group("/members")
	{
		router.POST("/register", controller.RegisterMember)
		router.POST("/login", controller.LoginMember)
		router.GET("/:memId/valid", controller.IsValidMember)
	}

	return engine
}
