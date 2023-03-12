package utils

import (
	. "StudentManger/module"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Sql = Init_gorm()

func Init_gorm() *gorm.DB {
	//配置MySQL连接参数
	username := "root"        //账号
	password := "root"        //密码
	host := "127.0.0.1"       //数据库地址，可以是Ip或者域名
	port := 3306              //数据库端口
	Dbname := "StudentManger" //数据库名

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	//类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	//连接MYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	LoadSqlTable(db)
	return db
}

func LoadSqlTable(db *gorm.DB) {
	err := db.AutoMigrate(&Student{})
	if err != nil {
		panic("Student表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&Teacher{})
	if err != nil {
		panic("Teacher表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&Admin{})
	if err != nil {
		panic("Admin表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&DormInfo{})
	if err != nil {
		panic("DormInfo表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&DormStaff{})
	if err != nil {
		panic("DormStaff表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&CourseInfo{})
	if err != nil {
		panic("CourseInfo表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&CourseStaff{})
	if err != nil {
		panic("CourseStaff表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&Score{})
	if err != nil {
		panic("Score表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&ConfirmInfo{})
	if err != nil {
		panic("Confirm表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&ConfirmResult{})
	if err != nil {
		panic("Teacher表同步失败, error=" + err.Error())
	}
	err = db.AutoMigrate(&AskLeave{})
	if err != nil {
		panic("AskLeave表同步失败, error=" + err.Error())
	}
}
