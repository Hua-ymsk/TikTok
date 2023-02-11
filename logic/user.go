package logic

import (
	"fmt"
	"github.com/jinzhu/copier"
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
func (logic *UserRegisterLogic) RegisterUser(user models.User) utils.Response {
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
	userNameCheck, _, _ := mysql.QueryUserName(user.UserName)
	if userNameCheck {
		return utils.CommonResponse(-1, "用户名存在", -1, "")
	}
	//注册用户
	//新增设置昵称为名称
	user.NickName = user.UserName
	userid := mysql.RegisterUser(&user)
	//生成token
	atoken, _, _ := middleware.GenToken(userid)
	return utils.CommonResponse(0, "用户注册成功", userid, atoken)
}

func (logic *UserLoginLogic) LoginUser(user models.User) utils.Response {
	//登录
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
	userNameCheck, User, _ := mysql.QueryUserName(user.UserName)
	if !userNameCheck {
		return utils.CommonResponse(-1, "用户名不存在或密码错误", -1, "")
	}
	//验证密码
	if user.PassWord != User.PassWord {
		return utils.CommonResponse(-1, "用户名不存在或密码错误", -1, "")
	}
	//生成token
	atoken, _, _ := middleware.GenToken(User.ID)

	return utils.CommonResponse(0, "用户登录成功", User.ID, atoken)
}

func (logic *UserInfoLogic) UserInfo(userid int64, id int64) utils.CResponse {
	var responseUser types.User
	//这里查询的是当前要查询的用户
	CheckUser, isfollow, err := mysql.QueryUserID(userid, id)
	if err != nil {
		return utils.CCResponse(-1, "查询失败", nil)
	}

	if err := copier.Copy(&responseUser, &CheckUser); err != nil {
		fmt.Println("copy err:", err)
		return utils.CCResponse(-1, "查询失败", nil)
	}
	responseUser.IsFollow = isfollow
	return utils.CCResponse(0, "用户信息获取成功", responseUser)
}
