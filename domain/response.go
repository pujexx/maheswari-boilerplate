package domain

type Response struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}
