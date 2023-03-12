package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddAskLeave(leave module.AskLeave) bool {
	if err := utils.Sql.Create(&leave).Error; err != nil {
		fmt.Println("AddAskLeave Error:", err)
		return false
	}
	return true
}

func FindAskLeaveByUid(uid string) module.AskLeave {
	var ask module.AskLeave
	err := utils.Sql.Where("uid = ?", uid).First(&ask)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return module.AskLeave{}
	}
	return ask
}

func DeleteAskLeaveByUid(uid string) bool {
	if err := utils.Sql.Where("uid = ?", uid).Delete(&module.AskLeave{}).Error; err != nil {
		return false
	}
	return true
}

func UpdateAskLeaveByUid(uid string, column string, value string) bool {
	if err := utils.Sql.Model(&module.AskLeave{}).Where("uid = ?", uid).Update(column, value).Error; err != nil {
		return false
	}
	return true
}
