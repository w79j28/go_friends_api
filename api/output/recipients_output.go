package output

//RecipientsOutput output
type RecipientsOutput struct {
	Output
	Recipients []string `json:"recipients"`
}
