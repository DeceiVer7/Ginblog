package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加收藏记录
func AddStarRecordAPI(c *gin.Context) {
	var data model.Star
	_ = c.ShouldBindJSON(&data)

	code := model.CreateStarRecord(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除收藏记录
func DeleteStarRecordAPI(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userid"))
	aid, _ := strconv.Atoi(c.Param("aid"))

	code := model.DeleteStarRecord(userid, aid)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询收藏记录
func GetStarRecordsAPI(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userid"))
	aid, _ := strconv.Atoi(c.Param("aid"))

	data, code := model.GetStarRecord(userid, aid)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
