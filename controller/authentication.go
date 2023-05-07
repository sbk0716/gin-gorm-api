package controller

import (
	"diary_api/helper"
	"diary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================
// Register function
// ============================================================
func Register(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		// Returns StatusBadRequest(400)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	// Executes User Save function.
	savedUser, err := user.Save()

	if err != nil {
		// Returns StatusBadRequest(400)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Returns StatusCreated(201)
	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

// ============================================================
// Login function
// ============================================================
func Login(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		// Returns StatusBadRequest(400)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Executes User FindUserByUsername function.
	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		// Returns StatusBadRequest(400)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Executes User ValidatePassword function.
	err = user.ValidatePassword(input.Password)

	if err != nil {
		// Returns StatusBadRequest(400)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Executes Helper GenerateJWT function.
	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		// Returns StatusBadRequest(400)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Returns StatusOK(200)
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
