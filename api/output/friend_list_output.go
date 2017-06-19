package output

//FriendListOutput output
type FriendListOutput struct {
	Output
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}
