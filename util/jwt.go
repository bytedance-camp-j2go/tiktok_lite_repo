package util

import (
	"errors"
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/dgrijalva/jwt-go/v4"
	// "github.com/golang-jwt/jwt"
	"time"
)

type CustomClaims struct {
	model.User
	jwt.StandardClaims
}

var MySecret = []byte("a1b2c3d4")

// GetToken 创建token，在调用此方法时，需要传入user对象
// 注意：这块传的user对象中不能存放密码等敏感信息，传入之前需删掉相关信息
func GetToken(user model.User) (string, error) {
	// 设置jwt的payload
	claim := CustomClaims{
		user,
		jwt.StandardClaims{}, // 存放签发者、签发时间、过期时间等信息，这块不设置过期时间
	}

	// 获取token，使用HS256加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// ParseToken 解析token，返回用户信息
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		fmt.Println("token解析错误", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效token")
}

// RefreshToken 刷新token，若在创建token时设置了过期时间，则需要在用户每次登录时刷新token
func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(time.Minute * 10))
		return GetToken(claims.User)
	}
	return "", errors.New("Cloudn't handle this token")
}
