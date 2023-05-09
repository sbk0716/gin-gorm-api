package controller

import (
	"diary_api/helper"
	"diary_api/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Register function:

1. Executes the validation.

2. Creates user model.

3. Executes User Save function.

4. If the Save function is successfully executed, StatusCreated(201) is returned.
*/
func Register(context *gin.Context) {
	var input model.AuthenticationInput
	// Sets the address of a variable(input).
	ptrInput := &input

	// Executes the validation.
	// FYI: https://gin-gonic.com/docs/examples/binding-and-validation/
	if err := context.ShouldBindJSON(ptrInput); err != nil {
		// If the validation fails, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// If the validation passes, the variable is filled with the request data.
	fmt.Printf("input: %#v\n", input)

	// Creates user model.
	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}
	fmt.Printf("user: %#v\n", user)

	// Sets the address of a variable(user).
	ptrUser := &user
	// Executes User Save function.
	// It returns the address of the pointer variable(ptrUser).
	savedUser, err := ptrUser.Save()
	fmt.Printf("savedUser: %#v\n", savedUser)

	if err != nil {
		// If the Save function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// If the Save function is successfully executed, StatusCreated(201) is returned.
	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

// ============================================================
// Login function
// ============================================================
/*
Login function:

1. Executes the validation.

2. Executes User FindUserByUsername function.

3. Executes User ValidatePassword function.

4. Executes Helper GenerateJWT function.

5. If the GenerateJWT function is successfully executed, StatusOK(200) is returned.
*/
func Login(context *gin.Context) {
	var input model.AuthenticationInput
	// Sets the address of a variable(input).
	ptrInput := &input

	// Executes the validation.
	// FYI: https://gin-gonic.com/docs/examples/binding-and-validation/
	if err := context.ShouldBindJSON(ptrInput); err != nil {
		// If the validation fails, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// If the validation passes, the variable is filled with the request data.
	fmt.Printf("input: %#v\n", input)

	// Executes User FindUserByUsername function.
	user, err := model.FindUserByUsername(input.Username)
	fmt.Printf("user: %#v\n", user)

	if err != nil {
		// If the FindUserByUsername function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Executes User ValidatePassword function.
	err = user.ValidatePassword(input.Password)

	if err != nil {
		// If the ValidatePassword function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Executes Helper GenerateJWT function.
	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		// If the GenerateJWT function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("jwt: %#v\n", jwt)

	// If the GenerateJWT function is successfully executed, StatusOK(200) is returned.
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
