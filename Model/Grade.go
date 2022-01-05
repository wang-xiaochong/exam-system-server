package model

import (
	database "Exam/Database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Grades = database.DB.Collection("grade")

//成绩接口
type Grade struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" `
	Major       string             `bson:"major" json:"major"` //专业
	Account     string             `bson:"account"    json:"account"    form:"account"`
	Grade       int                `bson:"grade" json:"grade"`             //分数
	SubjectName string             `bson:"subjectName" json:"subjectName"` //科目名称
}
