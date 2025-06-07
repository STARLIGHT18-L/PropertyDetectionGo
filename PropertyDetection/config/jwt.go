package config

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	TokenType    = "bearer"
	AppGuardName = "property-detection"
)

type JwtConfig struct {
}

var JwtService = new(JwtConfig)

// 所有需要颁发 token 的用户模型必须实现这个接口
type JwtUser interface {
	GetUid() string
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.StandardClaims
}
type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// CreateToken 生成 Token
func (config *JwtConfig) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + Boot.Config.Jwt.JwtTtl,
				Id:        user.GetUid(),
				Issuer:    GuardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)
	tokenStr, err := token.SignedString([]byte(Boot.Config.Jwt.Secret))
	tokenData = TokenOutPut{
		tokenStr,
		int(Boot.Config.Jwt.JwtTtl),
		TokenType,
	}
	return
}
