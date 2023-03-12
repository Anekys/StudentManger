package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddDormInfo(info module.DormInfo) bool {
	if err := utils.Sql.Create(&info).Error; err != nil {
		fmt.Println("AddDormInfo Error:", err)
		return false
	}
	return true
}

func FindDormInfoById(id string) module.DormInfo {
	var dorm module.DormInfo
	err := utils.Sql.Where("dorm_id = ?", id).First(&dorm)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return module.DormInfo{}
	}
	return dorm
}

func DeleteDormInfoById(id string) bool {
	if err := utils.Sql.Where("id = ?", id).Delete(&module.DormInfo{}).Error; err != nil {
		return false
	}
	return true
}
func UpdateDormInfoById(id string, column string, value string) bool {
	if err := utils.Sql.Model(&module.DormInfo{}).Where("id = ?", id).Update(column, value).Error; err != nil {
		return false
	}
	return true
}

func AddDormStaff(staff module.DormStaff) bool {
	if err := utils.Sql.Create(&staff).Error; err != nil {
		fmt.Println("AddDormStaff Error:", err)
		return false
	}
	return true
}

func FindDormStaffByDid(did string) module.DormStaff {
	var staff module.DormStaff
	err := utils.Sql.Where("dorm_id = ?", did).First(&staff)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return module.DormStaff{}
	}
	return staff
}

func DeleteDormStaffByDid(did string) bool {
	if err := utils.Sql.Where("dorm_id = ?", did).Delete(&module.DormStaff{}).Error; err != nil {
		return false
	}
	return true
}
