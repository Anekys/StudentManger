package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddCourse(info module.CourseInfo) bool {
	secretKey := []string{
		info.CName,
		info.TID,
	}
	kid := utils.CreateSecret(secretKey)
	info.KID = kid
	if err := utils.Sql.Create(&info).Error; err != nil {
		fmt.Println("AddCourse Error:", err)
		return false
	}
	return true
}

func FindCourseById(kid string) module.CourseInfo {
	var course module.CourseInfo
	err := utils.Sql.Where("kid = ?", kid).First(&course)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return module.CourseInfo{}
	}
	return course

}

func FindAllCourse(pageNum int) (courses []module.CourseInfo) {
	pageSize := 10
	utils.Sql.Where("").Order("kid ASC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&courses)
	return
}

func DeleteCourseById(kid string) bool {
	if err := utils.Sql.Where("kid = ?", kid).Delete(&module.CourseInfo{}).Error; err != nil {
		return false
	}
	return true
}

func UpdateCourseById(kid string, course module.CourseInfo) bool {
	if err := utils.Sql.Model(&module.CourseInfo{}).Where("kid = ?", kid).Updates(course).Error; err != nil {
		return false
	}
	return true
}

func AddCourseStudent(people module.CourseStaff) bool {
	if err := utils.Sql.Create(&people).Error; err != nil {
		fmt.Println("AddCourseStudent Error:", err)
		return false
	}
	return true
}

func FindCourseStudentByUID(uid string) module.CourseStaff {
	var people module.CourseStaff
	err := utils.Sql.Where("uid = ?", uid).First(&people)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return module.CourseStaff{}
	}
	return people
}

func DeleteCourseStudentByUid(uid string) bool {
	if err := utils.Sql.Where("uid = ?", uid).Delete(&module.CourseStaff{}).Error; err != nil {
		return false
	}
	return true
}
