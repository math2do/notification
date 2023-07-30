package dtos

type VerifyEmailReq struct {
	To   string `json:"to" binding:"required,email"`
	Link string `json:"link" binding:"required"`
}
