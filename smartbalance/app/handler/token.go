package handler

import (
	// "errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)


func generateJWT(username string) (string, error) {

	key, err := os.ReadFile("/application/gFunctions/cert/7a8e6b16-dec5-4500-873a-937f0d8e0c0a")

	log.Println("secret key", key)
	if err != nil {
		log.Println("create: parse key:", err)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	claims["aud"] = "intern"

	tokenString, err := token.SignedString(key)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}


func verifyJWT(tokenString string) (string, error){

	key, err := os.ReadFile("/application/gFunctions/cert/7a8e6b16-dec5-4500-873a-937f0d8e0c0a")

	log.Println("secret key", key)
	if err != nil {
		log.Println("create: parse key:", err)
	}


	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return key, nil
	})


	if err != nil {
		return "", nil
	}

	log.Println(token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", nil
	}
	return claims["sub"].(string), nil
}



// func getclaimsJWT(tokenString string) (string, error) {

// 	// var keyID string
// 	claims := jwt.MapClaims{}

// 	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

// 		// keyID, ok := token.Header["kid"].(string)
// 		log.Println(token.Header["alg"].(string))
// 		// log.Println(keyID)
// 		// if !ok {
// 		// 	return nil, errors.New("expecting JWT header to have string kid")
// 		// }

// 		key, err := os.ReadFile("/application/grpcFunc/cert/7a8e6b16-dec5-4500-873a-937f0d8e0c0a")
// 		// log.Println("pubKey", key)

// 		if err != nil {
// 			return "", fmt.Errorf("validate: parse key: %w", err)
// 		}

// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 		} else if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {

// 		} else {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return key, nil
// 	})



	// pubKey, err := os.ReadFile("/application/grpcFunc/cert/7a8e6b16-dec5-4500-873a-937f0d8e0c0a")

	// key, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	// if err != nil {
	// 	return "", fmt.Errorf("validate: parse key: %w", err)
	// }

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

	// 	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
	// 	} else if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {

	// 	} else {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	return key, nil
	// })

// 	if token == nil {
// 		log.Println("invalid token")
// 		return "Invalid Token", errors.New("token error: invalid signature")
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		log.Println("couldn't parse claims")
// 		return "Invalid Claims", errors.New("token error: claims")
// 	}

// 	aud := claims["aud"].(string)
// 	authorized := claims["authorized"].(bool)

// 	if aud == "intern" && authorized{
	
// 		return aud, nil
	
// 	} else if aud == "administrator" && authorized{
	
// 		return aud, nil
// 	}

// 	return "Failed to verify token", errors.New("token error: can't return ")
// }



