package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddAdmin(admin module.Admin) bool {
	if err := utils.Sql.Create(&admin).Error; err != nil {
		fmt.Println("AddAdmin Error:", err)
		return false
	}
	return true
}

func FindAdminByAid(aid string) module.Admin {
	var admin module.Admin
	err := utils.Sql.Where("aid = ?", aid).First(&admin)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return module.Admin{}
	}
	return admin
}

func FindAdminByEmailPassword(email string, password string) module.Admin {
	var admin module.Admin
	utils.Sql.Where("email = ? and password = ?", email, password).First(&admin)
	return admin
}

func DeleteAdminByAid(aid string) bool {
	if err := utils.Sql.Where("aid = ?", aid).Delete(&module.Admin{}).Error; err != nil {
		return false
	}
	return true
}

func UpdateAdminByAid(aid string, column string, value string) bool {
	if err := utils.Sql.Model(&module.ConfirmInfo{}).Where("aid = ?", aid).Update(column, value).Error; err != nil {
		return false
	}
	return true
}
