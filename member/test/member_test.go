package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todopoint/member/api"
	"todopoint/member/types"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {

	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/members", nil)
	mux := api.MakeWebHanlder()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []types.Member
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("hello", list[0].UserName)
	assert.Equal("hello2", list[1].UserName)
}
