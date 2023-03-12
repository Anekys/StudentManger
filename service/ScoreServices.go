package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddScore(score module.Score) bool {
	if err := utils.Sql.Create(&score).Error; err != nil {
		fmt.Println("AddScore Error:", err)
		return false
	}
	return true
}

func FindScoreByUid(uid string) module.Score {
	var score module.Score
	err := utils.Sql.Where("uid = ?", uid).First(&score)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return module.Score{}
	}
	return score
}

func DeleteScoreByUid(uid string) bool {
	if err := utils.Sql.Where("uid = ?", uid).Delete(&module.Score{}).Error; err != nil {
		return false
	}
	return true
}
func UpdateScoreByUid(uid string, column string, value string) bool {
	if err := utils.Sql.Model(&module.Score{}).Where("uid = ?", uid).Update(column, value).Error; err != nil {
		return false
	}
	return true
}
