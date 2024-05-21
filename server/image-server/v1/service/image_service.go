package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"modules/v2/common/httputils"
	"modules/v2/common/httputils/codes"
	"modules/v2/common/security/d7jwt"
	"path/filepath"
	"todopoint/image/data"
)

type ImageStore interface {
	Add(ctx *gin.Context, uid int, path string) error
	//Update(path string) error
}

type ImageService struct {
	store ImageStore
}

func NewImageService(store ImageStore) *ImageService {
	return &ImageService{
		store: store,
	}
}

func (i *ImageService) UploadImage(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError) {
	// Get user id
	uid, err := d7jwt.GetIdFromHeader(ctx)
	if err != nil {
		return nil, httputils.NewNetError(codes.Unauthorized, err)
	}

	// get file
	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, httputils.NewNetError(codes.InvalidFile, err)
	}

	// store in specific path
	uploadPath := "./uploads"
	filename := uuid.NewString() + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadPath, filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		return nil, httputils.NewNetError(codes.FileSaveFailed, err)
	}

	// save to mongo db
	// key, path, created_data
	err = i.store.Add(ctx, uid, filePath)
	if err != nil {
		return nil, httputils.NewNetError(codes.FileSaveFailed, err)
	}

	return httputils.NewSuccessBaseResponse(data.ImageUrl{Url: filePath}), nil

}
