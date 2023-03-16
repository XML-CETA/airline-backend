package generator

import (
	"main/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtGenerator struct {
}

func (generator *JwtGenerator) GenerateJwt(user *model.User) (string, error) {

	var secretKey = []byte(os.Getenv("SECRET_KEY"))

	claims := Claims{
		map[string]string{
			"username": user.Username,
			"role":     user.Role,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString(secretKey)

	return token, err
}
