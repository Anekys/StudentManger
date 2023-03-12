package module

type AskLeave struct {
	UID    string `gorm:"primary_key;column:uid"` // 请假学生的ID
	Cause  string `gorm:"column:cause"`           // 请假原因
	AID    string `gorm:"column:aid"`             // 处理的 管理员ID
	Status int    `gorm:"column:status"`          // 状态：0.待处理,1.批准2.拒绝
}
