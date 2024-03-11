package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todopoint/member/internal/controller"

	"todopoint/member/internal/model"
)

func TestJsonHandler(t *testing.T) {

	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/members", nil)
	mux := controller.MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []model.Member
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("hello", list[0].Username)
	assert.Equal("hello2", list[1].Username)
}
