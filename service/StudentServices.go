package service

import (
	"StudentManger/module"
	. "StudentManger/utils"
	"fmt"
)

func FindStudentByUid(Uid string) module.Student {
	var student module.Student
	Sql.Where("uid = ?", Uid).First(&student)
	return student
}
func FindStudentByEmailPassword(email string, password string) module.Student {
	var student module.Student
	Sql.Where("email = ? and password = ?", email, password).First(&student)
	return student
}
func FindAllStudents(pageNum int) (students []module.Student) {
	pageSize := 10
	Sql.Where("").Order("uid ASC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&students)
	return
}
func FindStudentsByClass(class string) (students []module.Student) {
	Sql.Model(&module.Student{}).Where("class = ?", class).Find(&students)
	return
}
func AddStudent(student module.Student) bool {
	if err := Sql.Create(&student).Error; err != nil {
		fmt.Println("StudentServices Error:", err)
		return false
	}
	return true
}

func DelStudentByUID(uid string) bool {
	if err := Sql.Where("uid = ?", uid).Delete(&module.Student{}).Error; err != nil {
		return false
	}
	return true
}

func UpdateStudentByUidWithKV(uid string, column string, value string) bool {
	if err := Sql.Model(&module.Student{}).Where("uid = ?", uid).Update(column, value).Error; err != nil {
		return false
	}
	return true
}

func UpdateStudentByUid(uid string, student module.Student) bool {
	if err := Sql.Model(&module.Student{}).Where("uid = ?", uid).Updates(student).Error; err != nil {
		return false
	}
	return true
}
