package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Author       string `gorm:"type:varchar(100);not null" json:"author"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	LikeCount    int    `gorm:"type:int;not null;default:0" json:"like_count"`
	StarCount    int    `gorm:"type:int;not null;default:0" json:"star_count"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64

	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateArtList).Error
	db.Model(&cateArtList).Where("cid =?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total
}

// GetArtInfo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err = db.Where("id = ?", id).Preload("Category").First(&art).Error
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// GetArt 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64

	err = db.Select("article.id, title, author, img, created_at, updated_at, `desc`, like_count, star_count, comment_count, read_count, category.name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total

}

// GetArtByAuthor 根据作者名查询文章列表
func GetArtByAuthor(pageSize int, pageNum int, author string) ([]Article, int, int64) {
	var articleListByAuthor []Article
	var err error
	var total int64

	err = db.Select("article.id, title, author, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").
		Limit(pageSize).Offset((pageNum-1)*pageSize).Order("created_at DESC").
		Joins("Category").Where("author = ?", author).
		Find(&articleListByAuthor).Error

	db.Model(&Article{}).Where("author = ?", author).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleListByAuthor, errmsg.SUCCSE, total
}

// GetArtByID 根据文章id查询文章列表
// func GetArtByID(pageSize int, pageNum int, id int) ([]Article, int, int64) {
// 	var articleListByID []Article
// 	var err error
// 	var total int64

// 	err = db.Select("article.id, title, author, img, created_at, updated_at, `desc`, comment_count, like_count,star_count").
// 		Limit(pageSize).Offset((pageNum-1)*pageSize).Order("created_at DESC").
// 		Where("id = ?", id).Find(&articleListByID).Error

// 	db.Model(&Article{}).Where("id = ?", id).Count(&total)

// 	if err != nil {
// 		return nil, errmsg.ERROR, 0
// 	}
// 	return articleListByID, errmsg.SUCCSE, total
// }

// SearchArticle 搜索文章标题
func SearchArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Select("article.id, title, author, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	//单独计数
	db.Model(&articleList).Where("title LIKE ?",
		title+"%",
	).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

// EditArt 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//点赞数量增加
func AddLikeCount(id int) int {
	err := db.Model(&Article{}).Where("id = ?", id).UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
	if err != nil {
		return errmsg.ERROR_LIKE_COUNT
	}
	return errmsg.SUCCSE
}

//点赞数量减少
func CancelLikeCount(id int) int {
	err := db.Model(&Article{}).Where("id = ?", id).UpdateColumn("like_count", gorm.Expr("like_count - ?", 1)).Error
	if err != nil {
		return errmsg.ERROR_CANCELLIKE_COUNT
	}
	if db.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//收藏数量增加
func AddStarCount(id int) int {
	err := db.Model(&Article{}).Where("id = ?", id).UpdateColumn("star_count", gorm.Expr("star_count + ?", 1)).Error
	if err != nil {
		return errmsg.ERROR_STAR_COUNT
	}
	return errmsg.SUCCSE
}

//收藏数量减少
func CancelStarCount(id int) int {
	err := db.Model(&Article{}).Where("id = ?", id).UpdateColumn("star_count", gorm.Expr("star_count - ?", 1)).Error
	if err != nil {
		return errmsg.ERROR_CANCELSTAR_COUNT
	}
	if db.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
