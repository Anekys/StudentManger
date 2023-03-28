package service

import (
	"StudentManger/module"
	"StudentManger/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// FindCourseAllStaffById 查询某个课程的人员以及其对应的分数,有分数就显示分数,没分数获得-1
func FindCourseAllStaffById(kid string, pageNum int) (staffs []module.ScoreStaff) {
	pageSize := 10
	utils.Sql.Table("course_staffs").
		Select("course_staffs.kid,course_staffs.uid, course_staffs.student, IFNULL(scores.score, -1) as score").
		Joins("LEFT JOIN scores ON course_staffs.uid = scores.uid AND course_staffs.kid = scores.kid").
		Where("course_staffs.kid = ?", kid).
		Order("course_staffs.uid ASC").Offset((pageNum - 1) * pageSize).Limit(pageSize).
		Scan(&staffs)
	//utils.Sql.Where("kid = ?", kid).Order("kid ASC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&staffs)
	return
}

// FindMineScore 根据uid查询所有课程的成绩
func FindMineScore(uid string) (scoreList []module.Score) {
	utils.Sql.Where("uid = ?", uid).Find(&scoreList)
	return
}

// FindScoreExist 通过kid和uid检查改名学生的成绩是否存在,存在返回真，不存在返回假
func FindScoreExist(kid string, uid string) bool {
	var res module.Score
	var empty module.Score
	utils.Sql.Where("kid = ? and uid = ?", kid, uid).Find(&res)
	if res == empty {
		return false
	}
	return true

}
func AddScore(score module.Score) bool {
	if err := utils.Sql.Create(&score).Error; err != nil {
		fmt.Println("AddScore Error:", err)
		return false
	}
	return true
}

func UpdatesScoreByKidAndUid(score module.Score) bool {
	err := utils.Sql.Where("kid = ? and uid = ?", score.KID, score.UID).Updates(score).Error
	if err != nil {
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
