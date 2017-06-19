package input

//SenderInput  input parameter
type SenderInput struct {
	Sender string `form:"sender" json:"sender" binding:"required,email"`
	Text   string `form:"text" json:"text" binding:"required"`
}
