package request

type CreateReqData struct {
	UserId   int    `json:"user_id"`
	BankName string `json:"bank_name"`
}
