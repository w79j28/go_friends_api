package input

type BlockInput struct {
	Requestor string `form:"requestor" json:"requestor" binding:"required,email"`
	Target    string `form:"target" json:"target" binding:"required,email"`
}
