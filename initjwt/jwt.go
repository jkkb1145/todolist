package initjwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"todo_list/conf"
	"todo_list/global"
	"todo_list/model"
)

var mySigningKey = []byte(conf.JWTKey)

type MyCustomClaims struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func CreateNewToken(user *model.User) (string, error) {
	claims := MyCustomClaims{
		user.UserID,
		user.UserName,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func PraseToken(token string) (*MyCustomClaims, error) {
	toke, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := toke.Claims.(*MyCustomClaims); ok {
		return claims, nil
	}
	return nil, errors.New("Failed to convert to pointer MyCustomClaims")
}

func StoreTokenInRedis(userID int, Token string) error {
	ctx := context.Background()
	userKey := fmt.Sprintf("user_tokens:%d", userID)

	err := global.RedisClient.HMSet(ctx, userKey,
		"token", Token,
	).Err()
	if err != nil {
		return err
	}

	return global.RedisClient.Expire(ctx, userKey, 2*24*time.Hour).Err()
}
