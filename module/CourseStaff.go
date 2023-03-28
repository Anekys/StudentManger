package module

type CourseStaff struct {
	ID      int64  `gorm:"primary_key;"`
	KID     string `gorm:"column:kid"`
	UID     string `gorm:"column:uid"`
	Student string `gorm:"column:student"`
}
