package main

import (
	"errors"
	"fmt"
	"gin-demo/basic/dao"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := dao.Connect(); err != nil {
		fmt.Println(errors.New("连接失败"))
		return
	}
	r := gin.Default()
	initRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
