package controller

import (
	"diary_api/helper"
	"diary_api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
AddEntry function:

1. Executes the validation.

2. Executes helper.CurrentUser function.

3. Sets userId to the value of the pointer variable(ptrInput).

4. Executes (*model.Entry).Save function.

5. If (*model.Entry).Save function is successfully executed, StatusCreated(201) is returned.
*/
func AddEntry(context *gin.Context) {
	var input model.Entry
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

	// Executes helper.CurrentUser function.
	// It returns the user struct.
	user, err := helper.CurrentUser(context)

	if err != nil {
		// If helper.CurrentUser function fails to execute, an error is returned..
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Sets userId to the value of the pointer variable(ptrInput).
	(*ptrInput).UserID = user.ID

	// Executes (*model.Entry).Save function.
	// It returns the address of the pointer variable(ptrInput).
	savedEntry, err := ptrInput.Save()
	fmt.Printf("savedEntry: %#v\n", savedEntry)

	if err != nil {
		// If (*model.Entry).Save function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// If (*model.Entry).Save function is successfully executed, StatusCreated(201) is returned.
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

/*
GetAllEntries function:

1. Executes helper.CurrentUser function.

2. If helper.CurrentUser function is successfully executed, StatusOK(200) is returned.
*/
func GetAllEntries(context *gin.Context) {
	// Executes helper.CurrentUser function.
	// It returns the user struct.
	user, err := helper.CurrentUser(context)

	if err != nil {
		// If helper.CurrentUser function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// If helper.CurrentUser function is successfully executed, StatusOK(200) is returned.
	context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

/*
GetEntry function:

1. Sets the entryId to a value extracted from the URL parameter.

2. Executes helper.CurrentUser function.
*/
func GetEntry(context *gin.Context) {
	// Sets the entryId to a value extracted from the URL parameter.
	entryId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		// If the URL parameter is invalid, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Printf("entryId: %#v\n", strconv.FormatUint(entryId, 10))

	// Executes helper.CurrentUser function.
	// It returns the user struct.
	user, err := helper.CurrentUser(context)

	if err != nil {
		// If helper.CurrentUser function fails to execute, StatusBadRequest(400) is returned.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result model.Entry
	// Extracs entries with matching entryId from the entry list of the target user.
	for _, entry := range user.Entries {
		if entry.ID == uint(entryId) {
			result = entry
			break
		}
	}

	fmt.Printf("result: %#v\n", result)

	if uint(entryId) != result.ID {
		errorMessage := fmt.Sprintf("The target entryId does not exist. [entryId: %s]", strconv.FormatUint(entryId, 10))
		// If entryId and result.ID do not match, StatusNotFound(404) is returned.
		context.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}

	// StatusOK(200) is returned.
	context.JSON(http.StatusOK, gin.H{"entry": result})
}
