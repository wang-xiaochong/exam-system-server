package model

import (
	database "Exam/Database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Subjects = database.DB.Collection("subject")
var TestStatus = database.DB.Collection("examStatus")

//科目
type Subject struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"      json:"_id,omitempty" `
	Name  string             `bson:"name" json:"name"`   //名称
	Major []string           `bson:"major" json:"major"` //专业
}

//考试情况
type ExamStatus struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"      json:"_id,omitempty" `
	SubjectName string             `bson:"subjectName" json:"subjectName"` //科目名称
	Student     []StuInfo          `bson:"student" json:"student"`
}
type StuInfo struct {
	Major       string `bson:"major" json:"major"`             //专业
	SubjectName string `bson:"subjectName" json:"subjectName"` //科目名称
	Account     string `bson:"account"    json:"account"    form:"account"`
	Name        string `bson:"name"       json:"name"       form:"name"`
	Status      string `bson:"status" json:"status"` //考试在线状态：0为离线 1为在线
}
