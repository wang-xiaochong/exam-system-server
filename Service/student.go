package service

import (
	model "Exam/Model"
	utils "Exam/Utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

//获取所有考试科目
func GetExamSubject(c *gin.Context, major string) []byte {
	cur, err := model.Subjects.Find(context.Background(), bson.M{"major": primitive.Regex{Pattern: major}})
	if err != nil {
		utils.Return(c, utils.EXCEPTION, err.Error())
		log.Println(err)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	var all []bson.M
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Println(err)
	}
_:
	cur.Close(context.Background())
	jsonStr, err := json.Marshal(all)
	return jsonStr
}

//获取考试科目内容
func GetExamSource(c *gin.Context, subject string) []byte {
	cur, err := model.Sources.Find(context.Background(), bson.M{"subjectName": subject})
	if err != nil {
		utils.Return(c, utils.EXCEPTION, err.Error())
		log.Println(err)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	var all []bson.M
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Println(err)
	}
_:
	cur.Close(context.Background())
	jsonStr, err := json.Marshal(all)
	return jsonStr
}

//保存答案
func SaveAnswer(c *gin.Context, answer model.Answer) {
	err := model.Answers.FindOne(context.Background(), bson.M{"account": answer.Account, "subjectName": answer.SubjectName})
	if err != nil {
		insertResult, err := model.Answers.InsertOne(context.Background(), answer)
		if err != nil {
			utils.Return(c, utils.EXCEPTION, err.Error())
			log.Println(err)
		}
		utils.Return(c, utils.SUCCESS, insertResult)
		return
	} else {
		update := bson.M{"$set": answer}
		updateResult, err := model.Answers.UpdateOne(context.Background(), bson.M{"account": answer.Account}, update)
		if err != nil {
			utils.Return(c, utils.EXCEPTION, err.Error())
			log.Println(err)
			return
		}
		utils.Return(c, utils.SUCCESS, updateResult)
		return
	}
}

//更改学生考试状态
func InExam(c *gin.Context, stu model.StuInfo) {
	stu.Status = "1"
	updateResult, err := model.TestStatus.UpdateOne(context.Background(), bson.M{"subjectName": stu.SubjectName},
		bson.M{"$push": bson.M{"student": stu}})
	if err != nil {
		utils.Return(c, utils.EXCEPTION, err.Error())
		log.Println(err)
		return
	}
	utils.Return(c, utils.SUCCESS, updateResult)
	return
}
func OutExam(c *gin.Context, stu model.StuInfo) {
	updateResult, err := model.TestStatus.UpdateOne(context.Background(), bson.M{"subjectName": stu.SubjectName, "student.account": stu.Account},
		bson.M{"$set": bson.M{"student.$.status": "0"}})
	if err != nil {
		utils.Return(c, utils.EXCEPTION, err.Error())
		log.Println(err)
		return
	}
	utils.Return(c, utils.SUCCESS, updateResult)
	return

}
