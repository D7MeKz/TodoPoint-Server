package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todopoint/member/api"
	"todopoint/member/internal/pkg/model"
)

func TestJsonHandler(t *testing.T) {

	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/members", nil)
	mux := api.MakeWebHanlder()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []model.Member
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("hello", list[0].UserName)
	assert.Equal("hello2", list[1].UserName)
}
