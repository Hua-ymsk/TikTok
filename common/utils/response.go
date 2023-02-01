package utils

type Response struct {
	StatusCode int64       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	UserId     int64       `json:"user_id"`
	Token      string      `json:"token"`
	Data       interface{} `json:"data,omitempty"`
}
type CResponse struct {
	StatusCode int64       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Data       interface{} `json:"user"`
}

func CommonResponse(code int64, message string, userid int64, token string) Response {
	return Response{
		StatusCode: code,
		StatusMsg:  message,
		UserId:     userid,
		Token:      token,
	}
}
func CCResponse(code int64, message string, data interface{}) CResponse {
	return CResponse{
		StatusCode: code,
		StatusMsg:  message,
		Data:       data,
	}
}
