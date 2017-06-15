// permission
package entity

type Permission struct {
	Requestor int64 `xorm:"pk"`
	Target    int64 `xorm:"pk"`
	Status    int   `xorm:"pk"`
}
