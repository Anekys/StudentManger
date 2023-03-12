package module

type ConfirmInfo struct {
	KID    string `gorm:"primary_key;column:kid"` // 课程ID
	TID    string `gorm:"column:tid"`             // 教师ID
	Secret string `gorm:"column:secret"`          // Redis key 考勤ID
}
