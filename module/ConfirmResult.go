package module

type ConfirmResult struct {
	KID    string `gorm:"primary_key;column:kid"` // 课程ID
	Secret string `gorm:"column:secret"`          // 考勤ID
	UID    string `gorm:"column:uid"`             // 学生ID
}
