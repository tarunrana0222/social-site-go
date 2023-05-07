package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tarunrana0222/social-site-go/config"
	"github.com/tarunrana0222/social-site-go/models"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	UserId string
	jwt.StandardClaims
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_Secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return "", nil
	}
}

func CreateToken(user models.User) (string, error) {
	userClaims := UserClaims{
		UserId: user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims).SignedString([]byte(config.JWT_Secret))
}

func CreatePasswordHash(password string) (string, error) {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(passBytes), err
}

func VerifyPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
