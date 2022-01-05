package model

import (
	database "Exam/Database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Answers = database.GetMongoDB().Collection("answer")

//答题答案
type Answer struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"      json:"_id,omitempty" `
	Major       string             `bson:"major" json:"major"` //专业
	Account     string             `bson:"account"    json:"account"    form:"account"`
	SubjectName string             `bson:"subjectName" json:"subjectName"` //科目名称
	StuAnswer   []ChoiceAndShort   `bson:"stuAnswer" json:"stuAnswer"`
}

type ChoiceAndShort struct {
	Id     int    `bson:"id" json:"id"`
	Answer string `bson:"answer" json:"answer"`
}
