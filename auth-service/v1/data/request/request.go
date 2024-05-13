package request

// RegisterRequest
// 모든 입력값은 필수이며, 비밀번호는 최소 8자 이상
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Username string `json:"username" binding:"required"`
}
