package entity

type VPermission struct {
	Requestor      int64
	Requestoremail string
	Target         int64
	Targetemail    string
	Status         int
}
