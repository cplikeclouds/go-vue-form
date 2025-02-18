package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/global"
	"server/models"
	"server/utils/errmsg"
	"strconv"
)

func GetFormList(c *gin.Context) {
	//传入页面大小和页数
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, _ := models.GetApplicantList(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
	global.Log.Println(data)
}

//
//func EditForm(c *gin.Context) {
//	var applicant models.Applicant
//	_ = c.ShouldBind(&applicant)
//	code := errmsg.SUCCESS
//
//}
