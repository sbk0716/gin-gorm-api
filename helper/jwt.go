package helper

import (
	"diary_api/model"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

/*
GenerateJWT function:

1. Sets tokenTTL.

2. Sets claims.

3. Creates a new Token with the specified signing method and claims.

4. Creates and returns a complete, signed JWT.
*/
func GenerateJWT(user model.User) (string, error) {
	// Sets tokenTTL.
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	// Sets claims.
	// MapClaims is a claims type that uses the map[string]interface{} for JSON decoding.
	claims := jwt.MapClaims{
		"id":  user.ID,                                                      // the user’s id (id)
		"iat": time.Now().Unix(),                                            // the time at which the token was issued (iat)
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(), // the expiry date of the token (eat).
	}
	// Creates a new Token with the specified signing method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Creates and returns a complete, signed JWT.
	return token.SignedString(privateKey)
}

/*
ValidateJWT function:

1. Executes getToken function to get the parsed token.

2. Executes Type assertions.

3. If token.Claims is castable to type jwt.MapClaims and the token is valid, nil is returned.
*/
func ValidateJWT(context *gin.Context) error {
	// Executes getToken function to get the parsed token.
	token, err := getToken(context)

	if err != nil {
		// If getToken function fails to execute,
		// it returns an error.
		return err
	}

	// Executes Type assertions.
	// FYI: https://go.dev/tour/methods/15
	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		// If token.Claims is castable to type jwt.MapClaims and the token is valid, nil is returned.
		return nil
	}

	// Returns an error if the token is invalid.
	return errors.New("invalid token provided")
}

/*
CurrentUser function:
*/
func CurrentUser(context *gin.Context) (model.User, error) {
	err := ValidateJWT(context)
	if err != nil {
		return model.User{}, err
	}

	token, _ := getToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	// Executes User FindUserById function.
	user, err := model.FindUserById(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

/*
jwtParseKeyFunc function:

1. Executes Type assertions.

2. If token.Method is castable to type *jwt.SigningMethodHMAC, privateKey and nil are returned.

jwt.Parse keyFunc:
keyFunc will receive the parsed token and should return the cryptographic key for verifying the signature.
*/
func jwtParseKeyFunc(token *jwt.Token) (interface{}, error) {
	// Executes Type assertions.
	// FYI: https://go.dev/tour/methods/15
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// If token.Method is not castable to type *jwt.SigningMethodHMAC, nil is returned.
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	// If token.Method is castable to type *jwt.SigningMethodHMAC, privateKey and nil are returned.
	return privateKey, nil
}

/*
getToken function:

1. Executes getTokenFromRequest function to get a JWT string from the bearer token.

2. Parses, validates, verifies the signature and returns the parsed token.

3. Returns the parsed token.

Example parsing and validating a token using the HMAC signing method:
FYI: https://pkg.go.dev/github.com/golang-jwt/jwt/v4@v4.5.0#example-Parse-Hmac
*/
func getToken(context *gin.Context) (*jwt.Token, error) {
	// Executes getTokenFromRequest function to get a JWT string from the bearer token.
	tokenString := getTokenFromRequest(context)
	// Parses, validates, verifies the signature and returns the parsed token.
	token, err := jwt.Parse(tokenString, jwtParseKeyFunc)
	// Returns the parsed token.
	return token, err
}

/*
getTokenFromRequest function:

1. Retrieves the bearer token from the request.

2. Bearer tokens come in the format `bearer <JWT>`,

3. So the retrieved string is split and the JWT string is returned.
*/
func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
