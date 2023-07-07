package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"fmt"
)

func AddAskLeave(leave module.AskLeave) bool {
	if err := utils.Sql.Create(&leave).Error; err != nil {
		fmt.Println("AddAskLeave Error:", err)
		return false
	}
	return true
}

func FindAskLeaveByUid(uid string, pageNum int) (asklist []module.AskLeave) {
	pageSize := 10
	utils.Sql.Where("uid = ?", uid).Order("status ASC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&asklist)
	return
}
func FindAskLeaveByLid(lid string) (ask module.AskLeave) {
	utils.Sql.Where("lid = ?", lid).First(&ask)
	return
}

func FindAllAskLeave(pageNum int) (asklist []module.AskLeave) {
	pageSize := 10
	utils.Sql.Where("").Order("status ASC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&asklist)
	return
}

func DeleteAskLeaveByUid(uid string) bool {
	if err := utils.Sql.Where("uid = ?", uid).Delete(&module.AskLeave{}).Error; err != nil {
		return false
	}
	return true
}

func UpdateAskLeaveByLid(lid string, ask module.AskLeave) bool {
	if err := utils.Sql.Model(&module.AskLeave{}).Where("lid = ?", lid).Updates(ask).Error; err != nil {
		fmt.Println("UpdateAskLeaveByLid error:", err)
		return false
	}
	return true
}
