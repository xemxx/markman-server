package common

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"markman-server/tools/config"
)

// Claims .
type Claims struct {
	Username string `json:"username"`
	UID      int    `json:"id"`
	jwt.RegisteredClaims
}

// GenerateToken .
func GenerateToken(username string, uid int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(60 * 24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
		UID:      uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "markman",
		},
	})

	jwtSecret := []byte(config.Cfg.GetString("app.jwt_secret"))
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

// ParseToken .
func ParseToken(tokenString string) (*Claims, error) {
	jwtSecret := []byte(config.Cfg.GetString("app.jwt_secret"))
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
