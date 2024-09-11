package dto

type Response struct {
	Code    int16       `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}
