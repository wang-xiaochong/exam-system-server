package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//成绩接口
type Grade struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" `
	Name            string             `bson:"name" json:"name" form:"name"`
	Major           string             `bson:"major" json:"major"` //专业
	GradeAndSubject []GS               `bson:"grade" json:"grade"` //科目和分数
}
type GS struct {
	Grade       int    `bson:"grade" json:"grade"`
	SubjectName string `bson:"subjectName" json:"subjectName"` //科目名称
}
