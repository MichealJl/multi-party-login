package proto

type ReqMpQQLoginParams struct {
	Code string `json:"code" valid:"required"`
}
