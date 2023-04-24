package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Like struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Aid      int    `gorm:"type:int;not null" json:"aid"`
	Status   int    `gorm:"type:int;not null" json:"status"`
	Userid   int    `gorm:"type:int;not null" json:"userid"`
}

// 添加点赞详情记录
func CreateLikeRecord(data *Like) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR_LIKE
	}
	return errmsg.SUCCSE
}

// 删除点赞记录
func DeleteLikeRecord(userid int, aid int) int {
	var like Like
	err := db.Where("userid = ? AND aid = ?", userid, aid).Delete(&like).Error
	if err != nil {
		return errmsg.ERROR_CANCELLIKE
	}
	return errmsg.SUCCSE
}

// 查询点赞记录
func GetLikeRecord(userid int, aid int) (Like, int) {
	var like Like
	if err := db.Where("userid = ? AND aid = ?", userid, aid).Find(&like).Error; err != nil && err != gorm.ErrRecordNotFound {
		return like, errmsg.ERROR
	}
	return like, errmsg.SUCCSE
}

//根据用户ID查询like表中所有的文章id
func GetLikedArticleIDsByUserID(userID int) ([]int, error) {
	var likedArticles []Like
	var articleIDs []int

	err := db.Model(&Like{}).Where("userid = ? AND status = 1", userID).Find(&likedArticles).Error
	if err != nil {
		return nil, err
	}

	for _, like := range likedArticles {
		articleIDs = append(articleIDs, like.Aid)
	}

	return articleIDs, nil
}

//用GetLikedArticleIDsByUserID方法查询到的文章id去查询出所有的文章
func GetArticlesByIDList(pageSize int, pageNum int, articleIDs []int) ([]Article, int, int64) {
	var likearticles []Article
	var total int64

	err := db.Where("id IN (?)", articleIDs).Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&likearticles).Error
	// 单独计数
	db.Model(&likearticles).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_LIKE_LIST, 0
	}

	return likearticles, errmsg.SUCCSE, total
}
