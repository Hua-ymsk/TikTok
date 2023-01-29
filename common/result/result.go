package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func ResponseErr(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		StatusCode: -1,
		StatusMsg:  msg,
	})
	return
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "请求成功",
	})
}
