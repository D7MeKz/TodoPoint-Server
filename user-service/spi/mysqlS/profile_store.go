package mysqlS

import (
	"github.com/gin-gonic/gin"
	"modules/d7mysql/ent"
	"modules/d7mysql/ent/profile"
	"modules/d7mysql/ent/user"
)

type ProfileStore struct {
	client *ent.Client
}

func NewProfileStore(client *ent.Client) *ProfileStore {
	return &ProfileStore{client: client}
}

func (p *ProfileStore) Set(ctx *gin.Context, uid int, imgUrl string) error {
	_, err := p.client.Profile.Update().Where(profile.HasUserWith(user.ID(uid))).SetImgURL(imgUrl).Save(ctx)
	if err != nil {
		return err
	}
	return err
}

func (p *ProfileStore) Get(ctx *gin.Context, uid int) (*ent.Profile, error) {
	// get user profile
	cProfile, err := p.client.Profile.Query().Where(profile.HasUserWith(user.ID(uid))).Only(ctx)
	if err != nil {
		return nil, err
	}
	return cProfile, nil
}
