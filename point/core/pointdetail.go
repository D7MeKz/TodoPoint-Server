package core

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/db/ent"
)

type PointDetailStore interface {
	Create(ctx *gin.Context) (*ent.PointInfo, error)
}
