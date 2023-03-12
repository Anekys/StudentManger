package service

import (
	"StudentManger/module"
	. "StudentManger/utils"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func AddUser(user module.User) bool {
	if err := Sql.Create(&user).Error; err != nil {
		return false
	}
	return true
}
func DeleteUser(user module.User) bool {
	if err := Sql.Where("username = ?", user.Username).Delete(&module.User{}).Error; err != nil {
		return false
	}
	return true
}

func FindUserByUserNameFirst(user module.User) module.User {
	//查询并返回第一条数据
	//定义需要保存数据的struct变量
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	User := module.User{}
	result := Sql.Where("username = ?", user.Username).First(&User)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return module.User{}
	}
	return User
}

func FindUserByUidFirst(user module.User) module.User {
	User := module.User{}
	result := Sql.Where("uid = ?", user.UID).First(&User)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return module.User{}
	}
	return User
}

func UpdateUser(user module.User, column string, value string) bool {
	if err := Sql.Model(&module.User{}).Where("username = ?", user.Username).Update(column, value).Error; err != nil {
		return false
	}
	return true
}

func AddStudentUser(username string, password string) (bool, string) {
	timeStr := GetNowTimeStamp()
	timeInt64, _ := strconv.ParseInt(timeStr, 10, 64)
	uid := CreatUid(username, password, timeStr)
	var user = module.User{
		Username:   username,
		Password:   password,
		UID:        uid,
		CreateTime: timeInt64,
		Identity:   3,
	}
	emptyUser := module.User{}
	if FindUserByUidFirst(user) != emptyUser {
		return false, "用户已存在"
	}
	return AddUser(user), "注册成功"
}
