package service

import (
	model "Exam/Model"
	utils "Exam/Utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"log"
)

//发布考试
func StartExam(c *gin.Context, exam model.ExamStatus) {
	var one model.ExamStatus
	err := model.TestStatus.FindOne(context.Background(), bson.M{"subjectName": exam.SubjectName}).Decode(&one)
	if err != nil {
		var s model.Subject
		err := model.Subjects.FindOne(context.Background(), bson.M{"name": exam.SubjectName}).Decode(&s)
		if err == nil {
			var user []model.User
			var stu []model.StuInfo
			for _, i := range s.Major {
				jsonStr := utils.Retrieve(c, "user", "major", i)
				var users []model.User
			_:
				json.Unmarshal(jsonStr, &users)
				user = append(user, users...)
			}
			for _, a := range user {
				stu = append(stu, model.StuInfo{
					Major:       a.Major,
					SubjectName: exam.SubjectName,
					Account:     a.Account,
					Name:        a.Name,
					Status:      "0",
				})
			}
			exam.Student = stu
		}
		insertOneResult, err := model.TestStatus.InsertOne(context.Background(), exam)
		if err != nil {
			utils.Return(c, utils.EXCEPTION, err.Error())
			log.Println(err)
		}
		log.Println("collection.InsertOne: ", insertOneResult)
		utils.Return(c, utils.SUCCESS, insertOneResult)
		return
	}
	utils.Return(c, utils.HAS_EXIST, one)
	return
}

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

//学生成绩汇总
func GetAllGrade(c *gin.Context, account, major string) {
	if account == "" && major != "" {
		jsonStr := utils.Retrieve(c, "grade", "major", major)
		var grades []model.Grade
	_:
		json.Unmarshal(jsonStr, &grades)
		utils.Return(c, utils.SUCCESS, grades)
	} else {
		cur, err := model.Grades.Find(context.Background(), bson.M{"account": account, "major": major})
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
		var grades []model.Grade
	_:
		json.Unmarshal(jsonStr, &grades)
		utils.Return(c, utils.SUCCESS, grades)
	}
}
