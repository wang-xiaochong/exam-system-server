package service

import (
	model "Exam/Model"
	utils "Exam/Utils"
	"context"
	"encoding/json"
	"fmt"
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
func GetExamSource(c *gin.Context, subject string) {
	var one model.Exam
	err := model.Sources.FindOne(context.Background(), bson.M{"subjectName": subject}).Decode(&one)
	if err == nil {
		utils.Return(c, utils.SUCCESS, one)
		return
	}
}

//保存答案
func SaveAnswer(c *gin.Context, answer model.Answer) {
	var one interface{}
	err := model.Answers.FindOne(context.Background(), bson.M{"account": answer.Account, "subjectName": answer.SubjectName}).Decode(&one)
	if err != nil {
		insertResult, err := model.Answers.InsertOne(context.Background(), answer)
		if err != nil {
			utils.Return(c, utils.EXCEPTION, err.Error())
			log.Println(err)
		}
		var source model.Exam
		err = model.Sources.FindOne(context.Background(), bson.M{"subjectName": answer.SubjectName}).Decode(&source)
		if err == nil {
			var grade model.Grade
			for _, b := range source.Questions1 {
				for _, a := range answer.StuAnswer {
					if a.Id == b.Id {
						if a.Answer == b.Choice {
							grade.Grade += 3
						}
					}
				}
			}
			for _, b := range source.Questions2 {
				for _, a := range answer.StuAnswer {
					if a.Id == b.Id {
						if a.Answer == b.Answer {
							grade.Grade += 8
						}
					}
				}
			}
			grade.Account = answer.Account
			grade.Major = answer.Major
			grade.SubjectName = answer.SubjectName
			insertResult2, err := model.Grades.InsertOne(context.Background(), grade)
			if err != nil {
				utils.Return(c, utils.EXCEPTION, err.Error())
				log.Println(err)
			}
			fmt.Println("-----", insertResult2)
		}
		utils.Return(c, utils.SUCCESS, insertResult)
		return
	}
	utils.Return(c, utils.HAS_EXIST, one)
	return
}

//更改学生考试状态
//进入考试 在线状态
func InExam(c *gin.Context, stu model.StuInfo) {
	stu.Status = "1"
	updateResult, err := model.TestStatus.UpdateOne(context.Background(), bson.M{"subjectName": stu.SubjectName, "student.account": stu.Account},
		bson.M{"$set": bson.M{"student.$.status": "1"}})
	if err != nil {
		utils.Return(c, utils.EXCEPTION, err.Error())
		log.Println(err)
		return
	}
	utils.Return(c, utils.SUCCESS, updateResult)
	return
}

//离线状态
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
