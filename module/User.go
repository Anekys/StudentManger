package module

type User struct {
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	//这里golang定义的Username变量和MYSQL表字段username一样，他们的名字可以不一样。
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	UID        string `gorm:"primary_key;column:uid"` // 用户唯一识别码
	Identity   int    `gorm:"column:identity"`        // 用户身份，枚举 1.管理员,2.老师,3.学生
	CreateTime int64  `gorm:"column:createtime"`
}
