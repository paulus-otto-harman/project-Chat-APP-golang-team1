package model

type Response struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}
type OTP struct {
	OTP string
}
type Token struct {
	Token string
}
