package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTwithID(userID int64) (string, error) {
	// Implement a simple JWT generation mechanism (for demonstration purposes only)
	// In production, use a library like github.com/dgrijalva/jwt-go
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateJWTwithIDUsernameName(userID int64, username, name string) (string, error) {
	// Implement a simple JWT generation mechanism (for demonstration purposes only)
	// In production, use a library like github.com/dgrijalva/jwt-go
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"name":     name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	log.Printf("Generating JWT for userID: %d, username: %s, name: %s", userID, username, name)
	log.Printf("JWT Token before signing: %v", token)
	log.Printf("JWT Claims: %v", token.Claims)
	log.Printf("JWT Claims user_id type is: %T", token.Claims.(jwt.MapClaims)["user_id"])
	log.Printf("JWT Claims username type is: %T", token.Claims.(jwt.MapClaims)["username"])
	log.Printf("JWT Claims name type is: %T", token.Claims.(jwt.MapClaims)["name"])
	log.Printf("JWT Signing Method: %v", token.Method)
	log.Printf("JWT Secret: %s", os.Getenv("JWT_SECRET"))
	log.Printf("Signing the token...")
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(tokenString, JwtSecret string) (int64, string, string, error) {
	// log.Printf("Inside of ParseJWT function")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// log.Printf("Parsing JWT token")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// log.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(JwtSecret), nil
	})
	// log.Printf("Token after parsing: %v", token)
	// log.Printf("error after parsing: %v", err)
	if err != nil {
		return 0, "", "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// log.Printf("Inside of logic to claim fields from jwt Token")
		// exp, foundExp := claims["exp"].(float64)
		// log.Printf("Found exp is %t and exp is: %f", foundExp, exp)
		userID := int64(claims["user_id"].(float64))
		_, foundUserId := claims["user_id"].(float64)
		// log.Printf("user_id type is %T: and value is: %d", claims["user_id"], int64(claims["user_id"].(float64)))
		// log.Printf("FoundUserID is %t and userID is: %d", foundUserId, userID)
		username, foundUsername := claims["username"].(string)
		// log.Printf("FoundUsername is %t and username is: %s", foundUsername, username)
		name, foundName := claims["name"].(string)
		// log.Printf("FoundName is %t and name is: %s", foundName, name)

		if foundUserId && foundUsername && foundName {
			// log.Printf("UserID from token claims: %d", userID)
			// log.Printf("Username from token claims: %s", username)
			// log.Printf("Name from token claims: %s", name)
			return userID, username, name, nil
		} else {
			log.Printf("One or more claims not found or of incorrect type")
			return 0, "", "", jwt.ErrInvalidKey
		}
		// if userID, ok := claims["user_id"].(int64); ok {
		// 	return userID, nil
		// }
	}
	return 0, "", "", jwt.ErrInvalidKey
}
