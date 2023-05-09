package model

// ============================================================
// Declare AuthenticationInput model
// ============================================================
// FYI: https://gin-gonic.com/docs/examples/binding-and-validation/
type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
