package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
	"todopoint/member/data/request"
	"todopoint/member/service"
	"todopoint/member/service/mocks"
)

func fixtures() (s *service.MemberService, m *mocks.MemberStore) {
	m = &mocks.MemberStore{}
	s = &service.MemberService{
		Store: m,
	}
	return s, m
}

func TestSuccess(t *testing.T) {
	s, m := fixtures()
	member1 := request.RegisterReq{
		Email:    "sample@naver.com",
		Password: "sample!",
	}

	userMatcher := func(u *request.RegisterReq) bool {
		return u.Email == member1.Email && u.Password == member1.Password
	}
	m.On("Create", mock.Anything, mock.MatchedBy(userMatcher)).Return(nil)

	// when
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	_, err := s.Store.Create(c, member1)

	// then
	assert.NoError(t, err.Err)
	m.AssertCalled(t, "Create", mock.Anything, mock.MatchedBy(userMatcher))

}
