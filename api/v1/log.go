package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 添加日志记录
func AddLogAPI(c *gin.Context) {
	var data model.Log
	_ = c.ShouldBindJSON(&data)

	code := model.CreateLog(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询所有日志记录
func GetAllLogsAPI(c *gin.Context) {

	data, code := model.GetAllLogs()

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除所有日志记录
func DeleteAllLogsAPI(c *gin.Context) {
	code := model.DeleteAllLogs()

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
