package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/models"
)

func HandleFormSubmit(c *gin.Context) {
	var data models.Applicant
	_ = c.ShouldBindJSON(&data)
	models.CreateApplicant(&data)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
		"msg":  "success",
	})

}
