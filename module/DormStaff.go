package module

type DormStaff struct {
	DormId string `gorm:"primary_key;column:dorm_id"`
	Name   string `gorm:"column:name"`
	Class  string `gorm:"column:class"`
	UID    string `gorm:"column:uid"`
}
