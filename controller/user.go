package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/logic"
)

type UserAPI struct{}

func (api *UserAPI) Register(c *gin.Context) {
	var logicregister logic.UserRegisterLogic
	err := c.ShouldBind(&logicregister)
	if err != nil {
		return
	}
	response := logicregister.RegisterUser(c)
	c.JSON(http.StatusOK, response)
}

func (api *UserAPI) Login(c *gin.Context) {
	var logiclogin logic.UserLoginLogic
	err := c.ShouldBind(&logiclogin)
	if err != nil {
		return
	}
	response := logiclogin.LoginUser(c)
	c.JSON(http.StatusOK, response)
}
func (api *UserAPI) UserInfo(c *gin.Context) {
	var logicuserinfo logic.UserInfoLogic
	err := c.ShouldBind(&logicuserinfo)
	if err != nil {
		return
	}
	response := logicuserinfo.UserInfo(c)
	c.JSON(http.StatusOK, response)
}
