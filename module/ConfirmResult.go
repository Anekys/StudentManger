package module

type ConfirmResult struct {
	ID      int64  `gorm:"primary_key;"`
	KID     string `gorm:"column:kid"`     // 课程ID
	CName   string `gorm:"column:name"`    // 课程名称
	Count   int64  `gorm:"column:count"`   // 本次考勤缺勤人数
	UID     string `gorm:"column:uid"`     // 学生ID
	Student string `gorm:"column:student"` // 学生姓名
	Time    string `gorm:"column:time"`    // 结束考勤时间
}
