package auth

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var secret string

func init() {
	//os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	secret = os.Getenv("ACCESS_SECRET")
	// For Test only
	if secret == "" {
		secret = "jdnfksdmfksd"
	}
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	//tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(tokenString string) (string, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		return "", err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return "", err
	}
	mp := token.Claims.(jwt.MapClaims)
	clientid := mp["clientid"]
	clid := clientid.(string)
	return clid, nil
}

// use this to create token
func CreateToken(userId string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["clientid"] = userId
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
