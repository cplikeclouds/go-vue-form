package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"server/global"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Config.Server.JwtKey),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// jwt中间件
// JWT 中间件
func JWTAuth() gin.HandlerFunc {
	JWT := NewJWT()
	JWT.JwtKey = []byte(global.Config.Server.JwtKey)
	if JWT.JwtKey == nil {
		panic("jwt_key 为空")
	}
	return func(c *gin.Context) {
		// 从请求头中提取 Token
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "未提供 Token",
			})
			c.Abort()
			return
		}

		// 检查 Token 格式（Bearer <token>）
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token 格式错误",
			})
			c.Abort()
			return
		}

		// 解析并验证 Token
		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JWT.JwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token 无效",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		// 提取 Claims
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			// 将用户信息附加到请求上下文中
			c.Set("user_id", claims.Username)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token 无效",
			})
			c.Abort()
		}
	}
}
