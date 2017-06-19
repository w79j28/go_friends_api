package input

//FriendInput input parameter
type FriendInput struct {
	Friends []string `description:"friends desc" json:"friends,required" binding:"required,dive,email"`
}
