package module

type Student struct {
	UID      string `gorm:"primary_key;column:uid"` // 学生ID
	Name     string `gorm:"column:name"`            //姓名
	Age      int    `gorm:"column:age"`             //年龄
	Gender   string `gorm:"column:gender"`          //性别 0.女 1.男
	Class    string `gorm:"column:class"`           //班级
	Phone    string `gorm:"column:phone"`           //联系方式
	Email    string `gorm:"column:email"`           //用户名
	PassWord string `gorm:"column:password"`        //密码
}
