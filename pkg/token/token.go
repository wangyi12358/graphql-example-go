package token

import (
	"github.com/golang-jwt/jwt/v5"
	"go-gin-example/pkg/config"
	"time"
)

type UserClaims struct {
	id int64
	jwt.RegisteredClaims
}

func GenerateToken(id int64) (*string, error) {
	claims := UserClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Config.Jwt.Expire)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                               // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                               // 生效时间
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(config.Config.Jwt.Secret))
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func ParseToken(token string) (*int64, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Jwt.Secret), nil
	})

	if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {

		return &claims.id, nil
	} else {
		return nil, err
	}
}
