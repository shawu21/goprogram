package token

import (
	"Program/constants"
	"Program/helper"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var MySecret = []byte("luxuetao")

// 生成JWT
func GenToken(username string) (string, error) {
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constants.TokenExpireDuration).Unix(),
			Issuer:    "goprogram",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

//解析JWT
func ParseToken(tokenString string) *MyClaims {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		panic("token error")
	}

	return claims
}

// JWT中间件
func JWTAuthMiddleware(c *gin.Context) {

	token := c.GetHeader("token")

	// 判断token是否存在
	if token == "" {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "TOKEN不存在或已失效!", ""))
		return
	}

	c.Set("nickname", ParseToken(token).Username)

	c.Next()
}

// JWT中间件
// func JWTAuthMiddleware() func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		authHeader := c.Request.Header.Get("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "请求头中的auth为空", ""))
// 			c.Abort()
// 			return
// 		}

// 		parts := strings.SplitN(authHeader, " ", 2)
// 		if !(len(parts) == 2 && parts[0] == "Bearer") {
// 			c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "请求头中auth格式有误", ""))
// 			c.Abort()
// 			return
// 		}

// 		mc, err := ParseToken(parts[1])
// 		if err != nil {
// 			c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "无效的token", ""))
// 			c.Abort()
// 			return
// 		}

// 		c.Set("username", mc.Username)
// 		c.Set("password", mc.Password)
// 		c.Next()
// 	}
// }
