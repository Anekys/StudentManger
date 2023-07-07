package module

type CourseInfo struct {
	KID       string `gorm:"primary_key;column:kid"` // 课程ID
	CName     string `gorm:"column:name"`            // 课程名
	CAbstract string `gorm:"column:abstract"`        //课程简介
	TID       string `gorm:"column:tid"`             //教师ID
	TName     string `gorm:"column:teacher"`         //教师名
	TEmail    string `gorm:"column:temail"`          //教师邮箱
	Status    int    `gorm:"column:status"`          //考勤状态: 0为考勤1正在考勤
}
