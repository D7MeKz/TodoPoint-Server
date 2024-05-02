package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http/httptest"
	"testing"
	"todopoint/member/out/ent"
	"todopoint/member/service"
	mocks "todopoint/member/test/mocks"
	"todopoint/member/utils/config"
	"todopoint/member/utils/data"
)

var (
	ctx *gin.Context
)

type memberSuite struct {
	suite.Suite
	store   *mocks.MemberStore
	service service.MemberService
}

func (suite *memberSuite) SetupSuite() {
	fmt.Println("Member Service Suite setup")
}

// Suite에서 각 테스트 실행 전에 실행
func (suite *memberSuite) SetupTest() {
	m := &mocks.MemberStore{}
	suite.store = m
	suite.service = service.MemberService{
		Store: m,
	}

	// Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	ctx = c
}

func (suite *memberSuite) TestCreate() {
	// When
	testMember := data.RegisterReq{
		Email:    "sample@naver.com",
		Username: "HelloUsername",
		Password: "sample!",
	}
	emptyMember := ent.Member{}
	entErr := gin.Error{}
	// Specify that inside ...
	suite.store.On("GetMemberByEmail", ctx, testMember.Email).Return(&emptyMember, entErr)
	if ent.IsNotFound(entErr) {
		suite.store.On("Create", ctx, &testMember).Return(&emptyMember, nil)
	}

	w1 := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w1)

	result, err := suite.service.CreateMember(c, testMember)

	suite.Nil(emptyMember, "Member are not nil.")
	suite.NoError(err, "No Error When member is created.")
	suite.Equal(result.Password, emptyMember.Password, "Test and result are same.")
}

func TestMemberService(t *testing.T) {
	// Init DB
	client, err := config.NewEntClient("../")
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	if err != nil {
		log.Println("Fail to initialize client")
	}

	// set client
	config.SetClient(client)
	suite.Run(t, new(memberSuite))
}
