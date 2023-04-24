package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Star struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Aid      int    `gorm:"type:int;not null" json:"aid"`
	Status   int    `gorm:"type:int;not null" json:"status"`
	Userid   int    `gorm:"type:int;not null" json:"userid"`
}

// 添加收藏详情记录
func CreateStarRecord(data *Star) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR_STAR
	}
	return errmsg.SUCCSE
}

// 删除收藏记录
func DeleteStarRecord(userid int, aid int) int {
	var star Star
	err := db.Where("userid = ? AND aid = ?", userid, aid).Delete(&star).Error
	if err != nil {
		return errmsg.ERROR_CANCELSTAR
	}
	return errmsg.SUCCSE
}

// 查询收藏记录
func GetStarRecord(userid int, aid int) (Star, int) {
	var star Star
	if err := db.Where("userid = ? AND aid = ?", userid, aid).Find(&star).Error; err != nil && err != gorm.ErrRecordNotFound {
		return star, errmsg.ERROR
	}
	return star, errmsg.SUCCSE
}

//根据用户ID查询star表中所有的文章id
func GetStarArticleIDsByUserID(userID int) ([]int, error) {
	var starArticles []Star
	var articleIDs []int

	err := db.Model(&Star{}).Where("userid = ? AND status = 1", userID).Find(&starArticles).Error
	if err != nil {
		return nil, err
	}

	for _, star := range starArticles {
		articleIDs = append(articleIDs, star.Aid)
	}

	return articleIDs, nil
}

//用GetStarArticleIDsByUserID方法查询到的文章id去查询出所有的文章
func GetStarArticlesByIDList(pageSize int, pageNum int, articleIDs []int) ([]Article, int, int64) {
	var stararticles []Article
	var total int64

	err := db.Where("id IN (?)", articleIDs).Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&stararticles).Error
	// 单独计数
	db.Model(&stararticles).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_STAR_LIST, 0
	}

	return stararticles, errmsg.SUCCSE, total
}
