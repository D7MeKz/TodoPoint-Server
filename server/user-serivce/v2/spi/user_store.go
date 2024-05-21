package spi

import (
	"github.com/gin-gonic/gin"
	"modules/v2/d7mysql/ent"
	"modules/v2/d7mysql/ent/profile"
	"modules/v2/d7mysql/ent/user"
	"todopoint/user/v2/data"
)

type UserStore struct {
	client *ent.Client
}

func NewUserStore(client *ent.Client) *UserStore {
	return &UserStore{client: client}
}

func (m *UserStore) FindOne(ctx *gin.Context, uid int) (*data.Me, error) {
	// Get profile with uid
	p, err := m.client.Profile.Query().Where(profile.HasUserWith(user.ID(uid))).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &data.Me{
		Username: p.Username,
		ImgUrl:   p.ImgURL,
	}, nil
}

func (m *UserStore) Create(ctx *gin.Context, info *data.UserInfo) error {

	// Save username to profile
	u, err := m.client.User.Get(ctx, info.Uid)
	if err != nil {
		return err
	}

	_, err = m.client.Profile.Create().SetUsername(info.Username).SetUser(u).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserStore) Update(ctx *gin.Context, uid int, me data.Me) error {
	// Get profile with uid
	p, err := m.client.Profile.Query().Where(profile.HasUserWith(user.ID(uid))).Only(ctx)
	// if is not found create new one
	if ent.IsNotFound(err) {
		_, cerr := m.client.Profile.Create().SetUserID(uid).SetUsername(me.Username).SetImgURL(me.ImgUrl).Save(ctx)
		if cerr != nil {
			return cerr
		}
	} else if err != nil {
		return err
	} else {
		_, err = m.client.Profile.UpdateOne(p).SetUsername(me.Username).SetImgURL(me.ImgUrl).Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
