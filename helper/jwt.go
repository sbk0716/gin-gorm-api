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
	fmt.Printf("tokenTTL: %#v\n", tokenTTL)
	// Sets claims.
	// MapClaims is a claims type that uses the map[string]interface{} for JSON decoding.
	claims := jwt.MapClaims{
		"id":  user.ID,                                                      // the user’s id (id)
		"iat": time.Now().Unix(),                                            // the time at which the token was issued (iat)
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(), // the expiry date of the token (eat).
	}
	fmt.Printf("claims: %#v\n", claims)
	// Creates a new Token with the specified signing method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Printf("token: %#v\n", token)
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
	fmt.Printf("token: %#v\n", token)

	if err != nil {
		// If getToken function fails to execute,
		// it returns an error.
		return err
	}
	// Executes Type assertions.
	// FYI: https://go.dev/tour/methods/15
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Printf("claims: %#v\n", claims)
	fmt.Printf("ok: %#v\n", ok)

	if ok && token.Valid {
		// If token.Claims is castable to type jwt.MapClaims and the token is valid, nil is returned.
		return nil
	}

	// Returns an error if the token is invalid.
	return errors.New("invalid token provided")
}

/*
CurrentUser function:

1. Executes ValidateJWT function.

2. Executes getToken function to get the parsed token.

3. Extracts userId from claims.

4. Executes model.FindUserByIdPreloadEntries function with userId.

5. If model.FindUserByIdPreloadEntries function is successfully executed, it returns the user struct and nil.
*/
func CurrentUser(context *gin.Context) (model.User, error) {
	// Executes ValidateJWT function.
	err := ValidateJWT(context)
	if err != nil {
		// If getToken function fails to execute,
		// it returns the empty struct and an error.
		return model.User{}, err
	}

	// Executes getToken function to get the parsed token.
	token, _ := getToken(context)
	// Executes Type assertions.
	// FYI: https://go.dev/tour/methods/15
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Printf("claims: %#v\n", claims)
	fmt.Printf("ok: %#v\n", ok)

	// Executes Type assertions.
	// FYI: https://go.dev/tour/methods/15
	floadId, floadOk := claims["id"].(float64)
	fmt.Printf("floadId: %#v\n", floadId)
	fmt.Printf("floadOk: %#v\n", floadOk)

	// Extracts userId from claims.
	userId := uint(floadId)
	fmt.Printf("userId: %#v\n", userId)

	// Executes model.FindUserByIdPreloadEntries function with userId.
	user, err := model.FindUserByIdPreloadEntries(userId)
	if err != nil {
		// If model.FindUserByIdPreloadEntries function fails to execute,
		// it returns the empty struct and an error.
		return model.User{}, err
	}
	// If model.FindUserByIdPreloadEntries function is successfully executed,
	// it returns the user struct and nil.
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
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	fmt.Printf("method: %#v\n", method)
	fmt.Printf("ok: %#v\n", ok)
	if !ok {
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
