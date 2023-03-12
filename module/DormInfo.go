package module

type DormInfo struct {
	DormId    string `gorm:"primary_key;column:dorm_id"` // 宿舍id
	DormNum   string `gorm:"column:dorm_num"`            //  宿舍楼号
	DormCount string `gorm:"column:dorm_count"`          // 宿舍人数
}
