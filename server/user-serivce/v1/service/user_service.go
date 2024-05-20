package service

import (
	"github.com/gin-gonic/gin"
	"modules/common/resource/d7image"
	"modules/common/security/d7jwt"
	"modules/common/server/httpdata"
	"modules/common/server/httpdata/d7errors"
	"modules/common/server/httpdata/d7errors/codes"
	"modules/d7mysql/ent"
	"path/filepath"
	"strconv"
	"todopoint/user/data/response"
)

type ProfileStorer interface {
	Set(ctx *gin.Context, uid int, imgUrl string) error
	Get(ctx *gin.Context, uid int) (*ent.Profile, error)
}
type UserService struct {
	store ProfileStorer
}

func NewUserService(store ProfileStorer) *UserService {
	return &UserService{
		store: store,
	}
}

func (u *UserService) Me(ctx *gin.Context) (*httpdata.BaseResponse, *d7errors.NetError) {
	// Get uid
	uid, err := d7jwt.GetIdFromHeader(ctx)
	if err != nil {
		return nil, d7errors.NewNetError(codes.TokenInvalid, err)
	}

	// get user profile
	profile, err := u.store.Get(ctx, uid)
	if err != nil {
		return nil, d7errors.NewNetError(codes.UserNotFound, err)
	}

	// Change to response
	res := response.Profile{Username: profile.Username, ImgUrl: profile.ImgURL}
	return httpdata.NewBaseResponse(codes.ProfileSuccess, res), nil

}

func (u *UserService) Create(ctx *gin.Context) (*httpdata.BaseResponse, *d7errors.NetError) {
	// if image is exist
	if ctx.Request.MultipartForm != nil {
		file, err := ctx.FormFile("image")
		if err != nil {
			return nil, d7errors.NewNetError(codes.BadImageHeader, err)
		}

		// get uid
		uid, err := d7jwt.GetIdFromHeader(ctx)
		if err != nil {
			return nil, d7errors.NewNetError(codes.TokenInvalid, err)
		}

		// upload file
		dir, err := d7image.NewFilePath(strconv.Itoa(uid))
		if err != nil {
			return nil, d7errors.NewNetError(codes.FileUploadError, err)
		}

		filename := d7image.FilenameGenerator()
		filePath := filepath.Join(dir, filename)
		err = ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			return nil, d7errors.NewNetError(codes.FileUploadError, err)
		}
		// set img_url in db
		err = u.store.Set(ctx, uid, filePath)
		if err != nil {
			return nil, d7errors.NewNetError(codes.ProfileUpdateError, err)
		}
	}

	return httpdata.NewBaseResponse(codes.ProfileSuccess, nil), nil
}
