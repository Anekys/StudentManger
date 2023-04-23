package module

// todo - 构建学生查询请假，提交请假申请的界面。以及老师的对应界面
type AskLeave struct {
	LID   string `gorm:"primary_key;column:lid"` // 假条的ID
	UID   string `gorm:"column:uid"`             // 请假学生的ID
	Name  string `gorm:"column:name"`            // 请假学生的姓名
	Cause string `gorm:"column:cause"`           // 请假原因
	AID   string `gorm:"column:aid"`             // 处理的 管理员ID
	Reason string `gorm:"column:reason"`	// 处理原因
	Status int `gorm:"column:status"` // 状态：0.待处理,1.批准2.拒绝
}
