package model

import (
	"ginblog/utils/errmsg"
)

type Profile struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"type:varchar(20)" json:"username"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	Qqchat    string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat    string `gorm:"type:varchar(100)" json:"wechat"`
	Weibo     string `gorm:"type:varchar(200)" json:"weibo"`
	Bili      string `gorm:"type:varchar(200)" json:"bili"`
	Email     string `gorm:"type:varchar(200)" json:"email"`
	Img       string `gorm:"type:varchar(200)" json:"img"`
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)" json:"icp_record"`
}

//新增用户个人信息
func CreateProfile(data *Profile) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetProfile 获取个人信息设置
func GetProfile(id int) (Profile, int) {
	var profile Profile
	err = db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCSE
}

// CheckProfileUser 查询Profile表中用户是否存在
func CheckProfileUser(name string) (code int) {
	var profile Profile
	db.Select("username").Where("username = ?", name).First(&profile)
	return errmsg.SUCCSE
}

// UpdateProfile 更新个人信息设置
func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
