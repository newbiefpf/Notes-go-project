package tools

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("newbie") //配置文件中自己配置的
var userId *int

// Claims是一些用户信息状态和额外的jwt参数
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   int    `json:"user_id"`
	jwt.StandardClaims
}

// 根据用户的用户名和密码参数token
func GenerateToken(username, password string, id int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24).Unix()

	claims := Claims{
		Username: username,
		Password: password,
		UserId:   id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime, // 过期时间
			Issuer:    "newbie",   //指定发行人
		},
	}
	// 该方法内部生成签名字符串，再用于获取完整、已签名的token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息(进而获取其中的用户名和密码)
func ParseToken(tokenString string) (*Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
