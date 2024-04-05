package networking

import (
	"encoding/json"
	"net/http"
	"time"
	"todopoint/common/errorutils/codes"
)

type ReqClient struct {
	Client *http.Client
}

func NewReqClient() *ReqClient {
	s, _ := time.ParseDuration("60s")
	return &ReqClient{
		Client: &http.Client{
			Timeout: s,
		},
	}
}

func parseBody(response *http.Response) (*ExternalInfo, error) {
	statusCode := response.StatusCode

	// Convert to empty interface
	res := make(map[string]interface{})
	err := json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	_ = response.Body.Close()

	// Extract ExternalInfo from response
	exInfo := extractFrom(res, statusCode)
	return exInfo, nil
}

type ExternalInfo struct {
	Code   codes.WebCode
	Status int
}

func extractFrom(res map[string]interface{}, status int) *ExternalInfo {
	// Convert to WebCode
	code, ok := res["code"].(codes.WebCode)
	if !ok {
		return nil
	}
	return &ExternalInfo{
		Code:   code,
		Status: status,
	}
}
