package model

import (
	database "Exam/Database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Sources = database.DB.Collection("examSource")

//考试试题库
type Exam struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"      json:"_id,omitempty" `
	SubjectName string             `bson:"subjectName" json:"subjectName"` //科目名称
	Questions1  []Choice           `bson:"questions1" json:"questions1"`   //选择题类型
	Questions2  []ShortAnswer      `bson:"questions2" json:"questions2"`   //选择题类型
}

//选择题类型
type Choice struct {
	Id            int    `bson:"id" json:"id"`
	Info          string `bson:"info" json:"info"` //问题
	Choice        string `bson:"choice" json:"choice"`
	CorrectChoice string `bson:"correctChoice" json:"correctChoice"`
	Answer        struct {
		A string `bson:"A" json:"A"`
		B string `bson:"B" json:"B"`
		C string `bson:"C" json:"C"`
		D string `bson:"D" json:"D"`
	} `bson:"answer" json:"answer"`
}

//简答题类型
type ShortAnswer struct {
	Id     int    `bson:"id" json:"id"`
	Info   string `bson:"info" json:"info"` //问题
	Answer string `bson:"answer" json:"answer"`
}
