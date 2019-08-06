package authentication

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dinhnguyen138/poker-backend/db"
	"github.com/dinhnguyen138/poker-backend/models"
	"github.com/dinhnguyen138/poker-backend/settings"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	middleWare *jwtmiddleware.JWTMiddleware
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			publicKey:  getPublicKey(),
			middleWare: jwtmiddleware.New(jwtmiddleware.Options{
				ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
					return getPublicKey(), nil
				},
				SigningMethod: jwt.SigningMethodRS512}),
		}
	}

	return authBackendInstance
}
func (backend *JWTAuthenticationBackend) GenerateToken(userUUID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * tokenDuration).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   userUUID,
	})
	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}
func (backend *JWTAuthenticationBackend) Authenticate(user *models.LoginMsg) string {
	return db.AuthUser(user.UserName, user.Password)
}
func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

// func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
// 	redisConn := redis.Connect()
// 	return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims["exp"]))
// }

// func (backend *JWTAuthenticationBackend) IsInBlacklist(token string) bool {
// 	redisConn := redis.Connect()
// 	redisToken, _ := redisConn.GetValue(token)

// 	if redisToken == nil {
// 		return false
// 	}

// 	return true
// }

func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}
func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
