package module

type Score struct {
	ID      int64  `gorm:"primary_key"`
	UID     string `gorm:"column:uid"`     // 学生ID
	Student string `gorm:"column:student"` // 学生名称
	TID     string `gorm:"column:tid"`     // 教师id
	Teacher string `gorm:"column:teacher"` // 教师名称
	KID     string `gorm:"column:kid"`     // 课程ID
	Course  string `gorm:"column:course"`  // 课程名
	Score   string `gorm:"column:score"`   // 成绩
}

// ScoreStaff 在成绩管理界面查询学生时的查询结果
type ScoreStaff struct {
	Kid     string
	UID     string
	Student string
	Score   int
}
