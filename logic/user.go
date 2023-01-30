package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"tiktok/common/utils"
	"tiktok/dao/mysql"
	"tiktok/middleware"
	"tiktok/models"
	"tiktok/types"
)

type UserRegisterLogic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserLoginLogic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserInfoLogic struct {
	Username string `json:"username"`
}

// 注册用户service
func (logic *UserRegisterLogic) RegisterUser(c *gin.Context) utils.Response {
	//获取注册数据
	user := models.User{
		UserName: c.Query("username"),
		PassWord: utils.Md5(c.Query("password")),
		NickName: c.Query("nickname"),
	}
	//检查用户名和密码
	if len(user.UserName) == 0 || len(user.PassWord) == 0 {
		return utils.CommonResponse(-1, "用户名或密码有误", -1, "")
	}
	if len(user.UserName) > 32 {
		return utils.CommonResponse(-1, "用名长度过长", -1, "")
	}
	if len(user.PassWord) > 32 {
		return utils.CommonResponse(-1, "密码过长", -1, "")
	}
	//查询当前用户名是否存在,如果存在返回，不存在继续操作
	userNameCheck, _, _, _, _, _, _ := mysql.QueryUserName(user.UserName)
	if userNameCheck {
		return utils.CommonResponse(-1, "用户名存在", -1, "")
	}
	//注册用户
	userid := mysql.RegisterUser(&user)
	//生成token
	atoken, _, _ := middleware.GenToken(userid)

	return utils.CommonResponse(1, "用户注册成功", userid, atoken)
}

func (logic *UserLoginLogic) LoginUser(c *gin.Context) utils.Response {
	//登录

	user := models.User{
		UserName: c.Query("username"),
		PassWord: utils.Md5(c.Query("password")),
	}
	//检查用户名和密码
	if len(user.UserName) == 0 || len(user.PassWord) == 0 {
		return utils.CommonResponse(-1, "用户名或密码有误", -1, "")
	}
	if len(user.UserName) > 32 {
		return utils.CommonResponse(-1, "用名长度过长", -1, "")
	}
	if len(user.PassWord) > 32 {
		return utils.CommonResponse(-1, "密码过长", -1, "")
	}
	//查询当前用户名是否存在,如果存在返回，不存在继续操作
	userNameCheck, userid, password, _, _, _, _ := mysql.QueryUserName(user.UserName)
	if !userNameCheck {
		return utils.CommonResponse(-1, "用户名不存在或密码错误", -1, "")
	}
	//验证密码
	if user.PassWord != password {
		return utils.CommonResponse(-1, "用户名不存在或密码错误", -1, "")
	}
	//生成token
	atoken, _, _ := middleware.GenToken(userid)

	return utils.CommonResponse(1, "用户登录成功", userid, atoken)
}

func (logic *UserInfoLogic) UserInfo(c *gin.Context) utils.CResponse {
	//获取用户信息
	userid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	//这里的user是我们要查询的user，不是自己
	user := models.User{
		ID: userid,
	}
	var responseUser types.User
	id, _ := c.Get("user_id")
	//这里查询的是当前要查询的用户
	fmt.Println("查询的用户是", id)
	_, responseUser.ID, _, responseUser.Name, responseUser.FollowerCount, responseUser.FollowCount, responseUser.IsFollow, _ = mysql.QueryUserID(user.ID, id)
	return utils.CCResponse(0, "用户信息获取成功", responseUser)
}
