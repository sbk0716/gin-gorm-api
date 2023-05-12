package model

/*
AuthenticationInput struct:

1. Username

2. Password

Model binding and validation:

To bind a request body into a type, use model binding.

FYI: https://gin-gonic.com/docs/examples/binding-and-validation/
*/
type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
