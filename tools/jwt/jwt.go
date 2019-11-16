package jwt

import (
	"markman-server/tools/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.Cfg.GetString("app.jwt_secret"))

//Claims .
type Claims struct {
	jwt.StandardClaims
	Username string      `json:"username"`
	UID      int         `json:"id"`
	_        interface{} //为了独立该工具封装做准备
}

//GenerateToken .
func GenerateToken(username string, uid int, d ...interface{}) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(365 * 24 * time.Hour)

	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "markman",
		},
		username,
		uid,
		d,
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
