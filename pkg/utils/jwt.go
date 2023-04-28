package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	key []byte
}

var (
	mySigningKey = []byte("woshishui")
)

// OasisClaims 自定义载荷
type OasisClaims struct {
	Username   string `json:"username"`
	BufferTime int64
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		key: mySigningKey,
	}
}

// NewOasisClaims 结构体实例化
func NewOasisClaims(name string) *OasisClaims {
	return &OasisClaims{
		Username: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "zhangshaodong",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)), // 测试用
		},
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(name string) (string, error) {
	// 使用指定的签名方法和声明创建一个新的Token, 加密方式后期替换
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewOasisClaims(name))
	// 加密key
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

// ParseToken 解析token
func (j *JWT) ParseToken(token string) (*OasisClaims, error) {
	// 解析
	tokenParse, err := jwt.ParseWithClaims(token, &OasisClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 解析token后返回
	if tokenParse != nil {
		if claims, ok := tokenParse.Claims.(*OasisClaims); ok && tokenParse.Valid {
			return claims, nil
		}
	}

	return nil, err
}
