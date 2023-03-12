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
	if err := Sql.Where("uid = ?", uid).Delete(&module.Student{}); err != nil {
		return false
	}
	return true
}

func UpdateStudentByUid(uid string, column string, value string) bool {
	if err := Sql.Model(&module.Student{}).Where("uid = ?", uid).Update(column, value).Error; err != nil {
		return false
	}
	return true
}
