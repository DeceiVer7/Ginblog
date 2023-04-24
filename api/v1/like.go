package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加点赞记录
func AddLikeRecordAPI(c *gin.Context) {
	var data model.Like
	_ = c.ShouldBindJSON(&data)

	code := model.CreateLikeRecord(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除点赞记录
func DeleteLikeRecordAPI(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userid"))
	aid, _ := strconv.Atoi(c.Param("aid"))

	code := model.DeleteLikeRecord(userid, aid)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询点赞记录
func GetLikeRecordsAPI(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userid"))
	aid, _ := strconv.Atoi(c.Param("aid"))

	data, code := model.GetLikeRecord(userid, aid)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
