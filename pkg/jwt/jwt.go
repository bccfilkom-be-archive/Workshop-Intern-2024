package jwt

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Interface interface {
	CreateJWTToken(userId uuid.UUID) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
}

type jsonWebToken struct {
	SecretKey   string
	ExpiredTime time.Duration
}

type Claims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func Init() Interface {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Fatalf("failed set expired time for jwt : %v", err.Error())
	}

	return &jsonWebToken{
		SecretKey:   secretKey,
		ExpiredTime: time.Duration(expiredTime) * time.Hour,
	}
}

func (j *jsonWebToken) CreateJWTToken(userId uuid.UUID) (string, error) {
	claim := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func (j *jsonWebToken) ValidateToken(tokenString string) (uuid.UUID, error) {
	var (
		claims Claims
		userId uuid.UUID
	)

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return userId, err
	}

	if !token.Valid {
		return userId, err
	}

	userId = claims.UserId

	return userId, nil
}

func (j *jsonWebToken) GetLoginUser(ctx *gin.Context) (entity.User, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return entity.User{}, errors.New("failed to get user")
	}

	return user.(entity.User), nil
}
