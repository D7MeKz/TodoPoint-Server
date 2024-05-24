package service_test

import (
	"modules/v2/common/httputils/codes"
	"modules/v2/common/testutils"
	"testing"
	"todopoint/user/v2/data"
	"todopoint/user/v2/service"
	"todopoint/user/v2/service/mocks"
)

var testCtx = testutils.GetTestGinContext()

func TestGetMe(t *testing.T) {
	// Create new mock store
	mockStore := new(mocks.Store)

	// Insert sample data
	uinfo := &data.UserInfo{Uid: 1, Username: "test"}
	mockStore.On("Create", testCtx, uinfo).Return(&data.Me{Username: "test", ImgUrl: ""}, nil)

	// Test
	tests := []struct {
		description string
		info        *data.UserInfo
		expected    *data.Me
		expectedErr codes.ErrorCode
	}{
		{
			description: "GetMeSuccess",
			info:        &data.UserInfo{Uid: 1, Username: "test"},
			expected:    &data.Me{Username: "test", ImgUrl: ""},
			expectedErr: 0,
		},
		{
			description: "Fail#1-User not found",
			info:        &data.UserInfo{Uid: 2, Username: "test"},
			expected:    nil,
			expectedErr: codes.NotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			mockStore.On("FindOne", testCtx, tc.info.Uid).Return(tc.expected, nil)
			svc := service.NewUserService(mockStore)
			res, err := svc.GetMe(testCtx, tc.info.Uid)

			// if error exist
			if err != nil {
				if err.Code != tc.expectedErr {
					t.Errorf("Expected %d, got %d", tc.expectedErr, err.Code)
					t.Logf("Error: %s", err.Error())
				}
			} else { // if error does not exist
				if res.Data != tc.expected {
					t.Errorf("Expected %v, got %v", tc.expected, res.Data)
				}
			}
		})
	}
}
