package module

type Score struct {
	UID   string  `gorm:"primary_key;column:uid"` // 学生ID
	TID   string  `gorm:"column:tid"`             // 教师id
	KID   string  `gorm:"column:kid"`             // 课程ID
	Score float32 `gorm:"column:score"`           // 成绩
}
