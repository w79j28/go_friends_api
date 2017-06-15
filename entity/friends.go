package entity

type Friends struct {
	Userid1 int64 `xorm:"pk"`
	Userid2 int64 `xorm:"pk"`
}
