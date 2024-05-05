package mysqlS

import (
	"errors"
	"github.com/gin-gonic/gin"
	"todopoint/auth/data/request"
	"todopoint/auth/spi/mysqlS/ent"
	"todopoint/auth/spi/mysqlS/ent/user"
	"todopoint/common/server/httpdata/auth"
)

type UserStore struct {
	client *ent.Client
}

func NewUserStore(client *ent.Client) *UserStore {
	return &UserStore{client: client}
}

func (m *UserStore) IsExist(ctx *gin.Context, data interface{}) (bool, error) {
	cred, ok := data.(auth.Credential)
	if !ok {
		return false, errors.New("IsExist : Invalid data type")
	}

	// Find user exist
	_, err := m.client.User.Query().Where(user.EmailEQ(cred.Email), user.PasswordEQ(cred.Password)).First(ctx)
	if ent.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *UserStore) Create(ctx *gin.Context, data interface{}) error {
	req, ok := data.(request.RegisterRequest)
	if !ok {
		return errors.New("Invalid data type")
	}

	_, err := m.client.User.Create().SetEmail(req.Email).SetPassword(req.Password).SetUsername(req.Username).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserStore) GetId(ctx *gin.Context, data interface{}) (int, error) {
	cred, ok := data.(auth.Credential)

	if !ok {
		return 0, errors.New("GetId : Invalid data type")
	}

	u, err := m.client.User.Query().Where(user.EmailEQ(cred.Email), user.PasswordEQ(cred.Password)).First(ctx)
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}
