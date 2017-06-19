package input

//EmailInput input parameter
type EmailInput struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}
