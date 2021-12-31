package model

import (
	database "Exam/Database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Answers = database.DB.Collection("answer")

//答题答案
type Answer struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"      json:"_id,omitempty" `
	Account     string             `bson:"account"    json:"account"    form:"account"`
	SubjectName string             `bson:"subjectName" json:"subjectName"` //科目名称
	Answers     []ChoiceAndShort   `bson:"answers" json:"answers"`
}
type ChoiceAndShort struct {
	Number int    `bson:"number" json:"number"` //题号
	Answer string `bson:"answer" json:"answer"` //答案
}
