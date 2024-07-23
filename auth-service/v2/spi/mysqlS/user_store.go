package mysqlS

import (
	"errors"
	"github.com/gin-gonic/gin"
	"modules/d7mysql/v2/ent"
	"modules/d7mysql/v2/ent/user"
	"todopoint/auth/v2/data"
)

type UserStore struct {
	client *ent.Client
}

func NewUserStore(client *ent.Client) *UserStore {
	return &UserStore{client: client}
}

func (m *UserStore) IsExist(ctx *gin.Context, d interface{}) (bool, error) {
	cred, ok := d.(data.Credential)
	if !ok {
		return false, errors.New("IsExist : Invalid httpdata type")
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

func (m *UserStore) Create(ctx *gin.Context, d interface{}) error {
	req, ok := d.(data.RegisterRequest)
	if !ok {
		return errors.New("Invalid httpdata type")
	}

	_, err := m.client.User.Create().SetEmail(req.Email).SetPassword(req.Password).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserStore) GetId(ctx *gin.Context, d interface{}) (int, error) {
	cred, ok := d.(data.Credential)

	if !ok {
		return 0, errors.New("GetId : Invalid httpdata type")
	}

	u, err := m.client.User.Query().Where(user.EmailEQ(cred.Email), user.PasswordEQ(cred.Password)).First(ctx)
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (m *UserStore) Modify(ctx *gin.Context, data interface{}) error {
	return nil
}

func (m *UserStore) IsValid(ctx *gin.Context, uid int) error {
	_, err := m.client.User.Query().Where(user.IDEQ(uid)).First(ctx)
	if err != nil {
		return err
	}
	return nil
}
