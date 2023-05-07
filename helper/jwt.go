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

// ============================================================
// GenerateJWT function
// ============================================================
func GenerateJWT(user model.User) (string, error) {
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
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

// ============================================================
// ValidateJWT function
// ============================================================
func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

// ============================================================
// CurrentUser function
// ============================================================
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

// ============================================================
// getToken function
// ============================================================
func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	// Parses, validates, verifies the signature and returns the parsed token.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

// ============================================================
// getTokenFromRequest function
// ============================================================
func getTokenFromRequest(context *gin.Context) string {
	// Retrieves the bearer token from the request.
	// Bearer tokens come in the format `bearer <JWT>`,
	// So the retrieved string is split and the JWT string is returned.
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
