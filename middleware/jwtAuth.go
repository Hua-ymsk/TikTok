package middleware

import (
	"errors"
	"fmt"
	"tiktok/common/result"
	"tiktok/setting"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims 内嵌jwt.StandardClaims,增加userid和username字段
type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

var (
	jwtConfig     = new(setting.JwtConfig)
	ISSUER        = jwtConfig.Issuer
	SECRET        = []byte("test")
	ACCESSEXPIRE  = jwtConfig.AccessExpire
	REFRESHEXPIRE = jwtConfig.RefreshExpire
)

// 定义错误
var (
	// 令牌过期
	ErrorTokenExpired = errors.New("token expired")
	ErrorInvalidToken = errors.New("invalid token")
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		mc, err := ParseToken(tokenStr)
		fmt.Println(mc.UserID)
		if err != nil {
			result.ResponseErr(c, "令牌无效或过期")
			c.Abort()
			return
		}
		// 验证签发人
		if mc.Issuer != ISSUER {
			result.ResponseErr(c, "令牌无效")
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("user_id", mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return SECRET, nil
}

// GenToken 生成access token 和 refresh token
func GenToken(userID int64) (aToken string, rToken string, err error) {
	// 创建一个自己的声明
	expireTime := time.Now().Add(2 * time.Hour)
	nowTime := time.Now()
	AccessClaims := MyClaims{
		UserID: userID,
		// JWT规定的官方字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间戳
			IssuedAt:  nowTime.Unix(),    //当前时间戳,
			Issuer:    ISSUER,
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, AccessClaims).SignedString(SECRET)

	// refresh token 不需要存任何自定义数据
	RefreshClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(viper.GetInt(REFRESHEXPIRE))).Unix(),
		Issuer:    ISSUER,
	}
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshClaims).SignedString(SECRET)
	return
}

// 解析tokenString
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	fmt.Println(token)
	fmt.Println(err)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		// 当AccessToken是过期错误
		if v.Errors == jwt.ValidationErrorExpired {
			err = ErrorTokenExpired
		}
		return
	}
	if !token.Valid {
		err = ErrorInvalidToken
	}
	return
}
