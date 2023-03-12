package module

type Teacher struct {
	Name     string `gorm:"column:name"`
	TID      string `gorm:"primary_key;column:tid"` // 教师ID
	Email    string `gorm:"column:email"`
	PassWord string `gorm:"column:password"`
}
