package auth

import (
	"markman-server/tools/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.Cfg.GetString("app.jwt_secret"))

//Claims .
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

//GenerateToken .
func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(365 * 24 * time.Hour)

	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "markman",
		},
		username,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

//ParseToken .
func ParseToken(token string) (*Claims, error) {
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
