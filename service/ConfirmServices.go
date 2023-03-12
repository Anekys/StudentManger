package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddConfirmInfo(info module.ConfirmInfo) bool {
	if err := utils.Sql.Create(&info).Error; err != nil {
		fmt.Println("AddConfirmInfo Error:", err)
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
