package controller

import (
	model "Exam/Model"
	service "Exam/Service"
	utils "Exam/Utils"
	"github.com/gin-gonic/gin"
)

//添加科目
func AddSubject(c *gin.Context) {
	var subject model.Subject
	if err := c.Bind(&subject); err != nil {
		utils.Return(c, utils.PARA_ERROR, err.Error())
		return
	}
	utils.CreateOne(c, "subject", subject, "name", subject.Name)
	return
}

//添加题库
func AddExamSource(c *gin.Context) {
	var sources model.Exam
	if err := c.Bind(&sources); err != nil {
		utils.Return(c, utils.PARA_ERROR, err.Error())
		return
	}
	utils.CreateOne(c, "examSource", sources, "subjectName", sources.SubjectName)
	return
}

//发布考试
func StartExam(c *gin.Context) {
	var exam model.ExamStatus
	if err := c.Bind(&exam); err != nil {
		utils.Return(c, utils.PARA_ERROR, err.Error())
		return
	}
	service.StartExam(c, exam)
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

//学生成绩汇总
func GetAllGrade(c *gin.Context) {
	account := c.Query("account")
	major := c.Query("major")
	service.GetAllGrade(c, account, major)
	return
}
