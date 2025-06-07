package middleware

import (
	"PropertyDetection/config"
	"PropertyDetection/tool"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			c.JSON(401, tool.TokenError())
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(config.TokenType)+1:]
		token, err := jwt.ParseWithClaims(tokenStr, &config.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Boot.Config.Jwt.Secret), nil
		})
		if err != nil {
			c.JSON(401, tool.TokenError())
			c.Abort()
			return
		}
		if !config.Boot.Cache.HasKey(tokenStr) { // 判断token是否存在缓存中
			c.JSON(401, tool.TokenError())
			c.Abort()
			return
		}
		claims := token.Claims.(*config.CustomClaims)
		if claims.Issuer != GuardName {
			c.JSON(401, tool.TokenError())
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Set("Id", claims.Id)
	}
}
