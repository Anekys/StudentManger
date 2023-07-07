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
	info.Status = 0
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

func FindChooseCourse(pageNum int, uid string) (courses []module.CourseInfo) {
	pageSize := 10
	utils.Sql.Table("course_infos").
		Joins("INNER JOIN course_staffs ON course_infos.kid = course_staffs.kid").
		Where("course_staffs.uid = ?", uid).Select("*").Scan(&courses).
		Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&courses)
	return
}

func FindNotChooseCourse(pageNum int, uid string) (courses []module.CourseInfo) {
	pageSize := 10
	utils.Sql.Table("course_infos").
		Select("course_infos.name,course_infos.kid,course_infos.abstract,course_infos.teacher").
		Joins("LEFT JOIN course_staffs ON course_infos.kid = course_staffs.kid AND course_staffs.uid = ?", uid).
		Where("course_staffs.kid IS NULL").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&courses)

	return
}

func FindCourseStaffByKid(kid string) (staffList []struct{}) {
	utils.Sql.Table("confirm_results").
		Select("confirm_results.uid, course_staffs.student, COUNT(confirm_results.uid) AS count").
		Joins("LEFT JOIN course_staffs.student ON confirm_results.uid = course_staffs.uid").
		Where("confirm_results.kid = ?", kid).
		Group("confirm_results.uid, course_staffs.student").
		Scan(&staffList)
	return
}

func ChooseCourse(courseStaff module.CourseStaff) bool {
	if err := utils.Sql.Create(&courseStaff).Error; err != nil {
		return false
	}
	return true
}

func RejectCourse(courseStaff module.CourseStaff) bool {
	if err := utils.Sql.Where("kid = ?", courseStaff.KID).Where("uid = ?", courseStaff.UID).Delete(&courseStaff).Error; err != nil {
		return false
	}
	return true
}

func FindTeacherCourse(tid string, pageNum int) (courses []module.CourseInfo) {
	pageSize := 10
	utils.Sql.Where("tid = ?", tid).Order("kid ASC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&courses)
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
		fmt.Println("UpdateCourseById error:", err)
		return false
	}
	return true
}

func UpdateCourseByIdWithField(kid string, field string, value interface{}) bool {
	err := utils.Sql.Model(&module.CourseInfo{}).Where("kid = ?", kid).Update(field, value).Error
	if err != nil {
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
