package service

import (
	model "Exam/Model"
	utils "Exam/Utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"log"
)

//获取考试状态
func GetExamStatus(c *gin.Context, subjectName string) {
	var one model.ExamStatus
	err := model.TestStatus.FindOne(context.Background(), bson.M{"subjectName": subjectName}).Decode(&one)
	if err != nil {
		utils.Return(c, utils.EXCEPTION, err.Error())
		log.Println(err)
		return
	} else {
		utils.Return(c, utils.SUCCESS, one)
		return
	}
}
