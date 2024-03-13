package router

import (
	"github.com/gin-gonic/gin"
	"todopoint/banking/controller"
)

func NewRouter(r *gin.Engine) {
	banking := r.Group("/banking")
	{
		banking.POST("/account/register", controller.RegisterAccount)
	}

}
