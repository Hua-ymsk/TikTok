package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/common/utils"
	"tiktok/logic"
	"tiktok/models"
)

type UserAPI struct{}

var salt string = "tiktok"

func (api *UserAPI) Register(c *gin.Context) {
	var logicregister logic.UserRegisterLogic
	user := models.User{
		UserName: c.Query("username"),
		PassWord: utils.Md5(c.Query("password"), salt),
		NickName: c.Query("nickname"),
	}
	response := logicregister.RegisterUser(user)
	c.JSON(http.StatusOK, response)
}

func (api *UserAPI) Login(c *gin.Context) {
	var logiclogin logic.UserLoginLogic
	//获取注册数据
	user := models.User{
		UserName: c.Query("username"),
		PassWord: utils.Md5(c.Query("password"), salt),
		NickName: c.Query("nickname"),
	}
	response := logiclogin.LoginUser(user)
	c.JSON(http.StatusOK, response)
}
func (api *UserAPI) UserInfo(c *gin.Context) {
	var logicuserinfo logic.UserInfoLogic
	userid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	id := c.GetInt64("user_id")
	fmt.Println(id)
	response := logicuserinfo.UserInfo(userid, id)
	c.JSON(http.StatusOK, response)
}
