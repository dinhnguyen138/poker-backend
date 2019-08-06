package services

import (
	"encoding/json"
	"net/http"

	"github.com/dinhnguyen138/poker-backend/core/authentication"
	"github.com/dinhnguyen138/poker-backend/models"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

func Login(requestUser *models.LoginMsg) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	userid := authBackend.Authenticate(requestUser)
	if userid != "" {
		token, err := authBackend.GenerateToken(userid)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(TokenAuthentication{token})
			return http.StatusOK, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(userid string) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(userid)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

// func Logout(req *http.Request) error {
// authBackend := authentication.InitJWTAuthenticationBackend()
// tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
// 	return authBackend.PublicKey, nil
// })
// if err != nil {
// 	return err
// }
// tokenString := req.Header.Get("Authorization")
// return authBackend.Logout(tokenString, tokenRequest)
// }
