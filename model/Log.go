package model

import (
	"ginblog/utils/errmsg"
)

type Log struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	Created_Time string `gorm:"type:varchar(100)" json:"created_time"`
	Content      string `gorm:"type:longtext" json:"content"`
}

// CreateLog 新增日志
func CreateLog(data *Log) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetAllLogs 查询日志列表
func GetAllLogs() ([]Log, int) {
	var logs []Log

	err := db.Select("id, created_time, content").Find(&logs).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return logs, errmsg.SUCCSE

}

// DeleteAllLogs 删除log表中所有数据
func DeleteAllLogs() int {
	err := db.Exec("DELETE FROM log").Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
