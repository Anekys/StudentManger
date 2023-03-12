package module

type CourseStaff struct {
	KID string `gorm:"primary_key;column:kid"`
	UID string `gorm:"column:uid"`
}
