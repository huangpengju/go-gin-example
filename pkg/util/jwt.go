package util

import (
	"go-gin-example/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 声明全局变量 jwtSecret 字节切片
var jwtSecret = []byte(setting.JwtSecret)

// 声明 Claims 要求结构
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Standard 标准
	// Claims 要求
	jwt.StandardClaims
}

// GenerateToken 生成 Token
func GenerateToken(username, password string) (string, error) {
	// 返回当前时间。
	nowTime := time.Now()
	// 到期时间
	// Add 返回时间点nowTime + 3 * time.Hour。(3小时)
	expireTime := nowTime.Add(3 * time.Hour)

	// 要求
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{ // 标准要求
			ExpiresAt: expireTime.Unix(), // 	设置到期时间（Unix 转换时间戳）
			// 设置发行人
			Issuer: "gin-blog",
		},
	}

	// 创建一个新 token ，采用签名方法
	// NewWithClaims(method SigningMethod, claims Claims)，method对应着SigningMethodHMAC struct{}，其包含SigningMethodHS256、SigningMethodHS384、SigningMethodHS512三种crypto.Hash方案
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 获取完整的 token
	// func (t *Token) SignedString(key interface{}) 该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析 Token
func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// func (m MapClaims) Valid() 验证基于时间的声明exp, iat, nbf，注意如果没有任何声明在令牌中，仍然会被认为是有效的。并且对于时区偏差没有计算方法
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
