package controller

import (
	model "Exam/Model"
	service "Exam/Service"
	utils "Exam/Utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//获取考试科目
func GetExamSubject(c *gin.Context) {
	major := c.Query("major")
	if major == "" {
		utils.Return(c, 500, "major为空")
		return
	} else {
		jsonStr := service.GetExamSubject(c, major)
		var subject []model.Subject
	_:
		json.Unmarshal(jsonStr, &subject)
		utils.Return(c, utils.SUCCESS, subject)
		return
	}
}

//获取考试科目内容
func GetExamSource(c *gin.Context) {
	subject := c.Query("subject")
	if subject == "" {
		utils.Return(c, 500, "subject为空")
		return
	} else {
		service.GetExamSource(c, subject)
		return
	}
}

//保存答案
func SaveAnswer(c *gin.Context) {
	var answer model.Answer
	if err := c.Bind(&answer); err != nil {
		utils.Return(c, utils.PARA_ERROR, err.Error())
		return
	}
	service.SaveAnswer(c, answer)
	return
}

//进入考试状态
func InExam(c *gin.Context) {
	var stu model.StuInfo
	if err := c.Bind(&stu); err != nil {
		utils.Return(c, utils.PARA_ERROR, err.Error())
		return
	}
	service.InExam(c, stu)
	return
}

//离开考试状态
func OutExam(c *gin.Context) {
	var stu model.StuInfo
	if err := c.Bind(&stu); err != nil {
		utils.Return(c, utils.PARA_ERROR, err.Error())
		return
	}
	service.OutExam(c, stu)
	return
}
