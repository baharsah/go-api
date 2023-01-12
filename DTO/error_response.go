package dto

type ErrorResult struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Validator interface{} `json:"validator"`
}
