package entity

// Permission table api_permission
type Permission struct {
	Requestor int64 `xorm:"pk"`
	Target    int64 `xorm:"pk"`
	Status    int   `xorm:"pk"`
}
