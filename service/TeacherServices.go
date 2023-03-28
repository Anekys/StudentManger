package service

import (
	"StudentManger/module"
	. "StudentManger/utils"
	"fmt"
)

func FindTeacherByTid(uid string) module.Teacher {
	var teacher module.Teacher
	Sql.Where("tid = ?", uid).First(&teacher)
	return teacher
}
func FindAllTeachers(pageNum int) (teachers []module.Teacher) {
	pageSize := 10
	Sql.Where("").Order("tid ASC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&teachers)
	return
}

func FindTeacherByEmailPassword(email string, password string) module.Teacher {
	var teacher module.Teacher
	Sql.Where("email = ? and password = ?", email, password).First(&teacher)
	return teacher
}

func AddTeacher(teacher module.Teacher) bool {
	uid := Md5Encrypt(teacher.Email)
	teacher.TID = uid
	if err := Sql.Create(&teacher).Error; err != nil {
		fmt.Println("TeacherServices Error:", err)
		return false
	}
	return true
}

func DeleteTeacher(tid string) bool {
	if err := Sql.Where("tid = ?", tid).Delete(&module.Teacher{}).Error; err != nil {
		return false
	}
	return true
}

func UpdateTeacherByTid(tid string, teacher module.Teacher) bool {
	if err := Sql.Model(&module.Teacher{}).Where("tid = ?", tid).Updates(teacher).Error; err != nil {
		return false
	}
	return true
}
