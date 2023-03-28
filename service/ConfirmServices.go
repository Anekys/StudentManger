package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func StartConfirm(kid string) (bool, string) {
	var staff []module.CourseStaff
	utils.Sql.Table("course_staffs").Select("*").Where("kid = ?", kid).Find(&staff)
	if len(staff) > 0 {
		staffs := utils.Staff2Map(staff)
		return utils.HMSet(kid, staffs), "未能将课程学院添加至考勤名单,请稍后重试或联系管理员!"
	} else {
		return false, "课程报名的人员不足,无法考勤!"
	}

}

func AddConfirmInfo(info module.ConfirmInfo) bool {
	if err := utils.Sql.Create(&info).Error; err != nil {
		fmt.Println("Add ConfirmInfo Error:", err)
		return false
	}
	return true
}

func FindTeacherConfirmResult(tid string) (results []module.ConfirmResult) {
	utils.Sql.Table("confirm_results").
		Joins("INNER JOIN course_infos ON confirm_results.kid = course_infos.kid").
		Where("course_infos.tid = ?", tid).
		Select("confirm_results.kid,confirm_results.time,confirm_results.name,confirm_results.count,COUNT(*) as countNum").
		Group("confirm_results.kid, confirm_results.time, confirm_results.name,confirm_results.count").Find(&results)
	return
}

func FindConfirmResultByKidAndTime(kid string, timeStr string) (results []module.ConfirmResult) {
	utils.Sql.Table("confirm_results").Where("kid = ? and time = ?", kid, timeStr).Find(&results)
	return
}

func AddConfirmResult(info interface{}) bool {
	switch v := info.(type) {
	case module.ConfirmResult:
		if err := utils.Sql.Create(&v).Error; err != nil {
			fmt.Println("Add ConfirmResult Error:", err)
			return false
		}
	case []module.ConfirmResult:
		if err := utils.Sql.Create(&v).Error; err != nil {
			fmt.Println("Add ConfirmResults Error:", err)
			return false
		}
	default:
		fmt.Println("Invalid argument type:", v)
		return false
	}
	return true
}

func FindConfirmInfoByKid(kid string) module.ConfirmInfo {
	var confirm module.ConfirmInfo
	err := utils.Sql.Where("kid = ?", kid).First(&confirm)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return module.ConfirmInfo{}
	}
	return confirm
}

func DeleteConfirmInfoByKid(kid string) bool {
	if err := utils.Sql.Where("kid = ?", kid).Delete(&module.ConfirmInfo{}).Error; err != nil {
		return false
	}
	return true
}

func UpdateConfirmInfoByKid(kid string, column string, value string) bool {
	if err := utils.Sql.Model(&module.ConfirmInfo{}).Where("kid = ?", kid).Update(column, value).Error; err != nil {
		return false
	}
	return true
}
