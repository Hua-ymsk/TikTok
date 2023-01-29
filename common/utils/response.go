package utils

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func CommonResponse(code int64, message string, userid int64, token string) Response {
	return Response{
		StatusCode: code,
		StatusMsg:  message,
		UserId:     userid,
		Token:      token,
	}
}
