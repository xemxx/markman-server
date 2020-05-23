package common

import (
	"markman-server/tools/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Claims .
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UID      int    `json:"id"`
}

//GenerateToken .
func GenerateToken(username string, uid int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(60 * 24 * time.Hour)

	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "markman",
		},
		username,
		uid,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(config.Cfg.GetString("app.jwt_secret"))
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

//ParseToken .
func ParseToken(token string) (*Claims, error) {
	jwtSecret := []byte(config.Cfg.GetString("app.jwt_secret"))
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
