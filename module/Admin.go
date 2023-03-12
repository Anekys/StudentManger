package module

type Admin struct {
	Email    string `gorm:"column:email"`           //用户名
	PassWord string `gorm:"column:password"`        //密码
	AID      string `gorm:"primary_key;column:aid"` //管理员ID
	Name     string `gorm:"column:name"`            //姓名
}
