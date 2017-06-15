package input

type EmailInput struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}
