package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//考试试题库
type Exam struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"      json:"_id,omitempty" `
	SubjectName string                 `bson:"subjectName" json:"subjectName"` //科目名称
	Questions1  map[string]Choice      `bson:"questions1" json:"questions1"`   //选择题类型
	Questions2  map[string]ShortAnswer `bson:"questions2" json:"questions2"`   //选择题类型
}

//选择题类型
type Choice struct {
	Question string `bson:"question" json:"question"` //问题
	A        string `bson:"a" json:"a"`
	B        string `bson:"b" json:"b"`
	C        string `bson:"c" json:"c"`
	D        string `bson:"d" json:"d"`
}

//简答题类型
type ShortAnswer struct {
	Question string `bson:"question" json:"question"` //问题
	Answer   string `bson:"answer" json:"answer"`
}
