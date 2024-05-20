package api

import (
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
)

type ImageOperator interface {
	UploadImage(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError)
	// DownloadImage(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError)
}

type ImageController struct {
	operator ImageOperator
}

func NewImageController(operator ImageOperator) *ImageController {
	return &ImageController{
		operator: operator,
	}
}

func (c *ImageController) Upload(ctx *gin.Context) {
	res, err := c.operator.UploadImage(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	res.OKSuccess(ctx)
}
