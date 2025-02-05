package helper

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nabinkatwal7/irlquest/model"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Hour * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func CurrentUser(context *gin.Context) (model.User, error){
	err := ValidateJWT(context)

	if err != nil {
		return model.User{}, err
	}

	token, _ := getToken(context)

	claims, _ := token.Claims.(jwt.MapClaims)

	userId := uint(claims["id"].(float64))

	user, err := model.FindUserById(userId)

	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}

func getToken(context *gin.Context)(*jwt.Token, error){
	tokenString, _ := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func getTokenFromRequest(context *gin.Context) (string, error) {
	bearerToken := context.Request.Header.Get("Authorization")
	if bearerToken == "" {
		return "", errors.New("not logged in: no authorization header")
	}

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", errors.New("not logged in: invalid token format")
	}

	return splitToken[1], nil
}