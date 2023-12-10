package jwtutil

import (
	"fmt"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func SubFromBearerToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", myerror.InvalidRequest
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return "", myerror.InvalidRequest // Bearerトークンが見つからない場合
	}

	token, err := ParseToken(tokenString)
	if err != nil {
		return "", myerror.InvalidRequest
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, ok := claims["sub"].(string)
		if !ok {
			return "", myerror.InvalidRequest // subフィールドが存在しない場合
		}
		return sub, nil
	}

	return "", myerror.InvalidRequest // その他のエラー
}
