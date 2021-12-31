package controller

import (
	model "Exam/Model"
	service "Exam/Service"
	utils "Exam/Utils"
	"github.com/gin-gonic/gin"
)

//发布考试
func StartExam(c *gin.Context) {
	var exam model.ExamStatus
	if err := c.Bind(&exam); err != nil {
		utils.Return(c, utils.PARA_ERROR, err.Error())
		return
	}
	utils.CreateOne(c, "examStatus", exam, "subjectName", exam.SubjectName)
	return
}

//获取在线考试状态
func FindStatus(c *gin.Context) {
	subjectName := c.Query("subjectName")
	if subjectName == "" {
		utils.Return(c, 500, "subjectName为空")
		return
	} else {
		service.GetExamStatus(c, subjectName)
		return
	}
}

