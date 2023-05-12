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

3. Executes (*model.User).Save function.

4. If (*model.User).Save function is successfully executed, StatusCreated(201) is returned.
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
	// Executes (*model.User).Save function.
	// It returns the address of the pointer variable(ptrUser).
	savedUser, err := ptrUser.Save()
	fmt.Printf("savedUser: %#v\n", savedUser)

	if err != nil {
		// If (*model.User).Save function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// If (*model.User).Save function is successfully executed, StatusCreated(201) is returned.
	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

/*
Login function:

1. Executes the validation.

2. Executes model.FindUserByUsername function.

3. Executes (*model.User).ValidatePassword function.

4. Executes helper.GenerateJWT function.

5. If helper.GenerateJWT function is successfully executed, StatusOK(200) is returned.
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

	// Executes model.FindUserByUsername function.
	user, err := model.FindUserByUsername(input.Username)
	fmt.Printf("user: %#v\n", user)

	if err != nil {
		// If model.FindUserByUsername function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Sets the address of a variable(user).
	ptrUser := &user
	// Executes (*model.User).ValidatePassword function.
	err = ptrUser.ValidatePassword(input.Password)

	if err != nil {
		// If (*model.User).ValidatePassword function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Executes helper.GenerateJWT function.
	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		// If helper.GenerateJWT function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("jwt: %#v\n", jwt)

	// If helper.GenerateJWT function is successfully executed, StatusOK(200) is returned.
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
