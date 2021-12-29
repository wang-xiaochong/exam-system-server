package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//科目
type Subject struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"      json:"_id,omitempty" `
	Name        string             `bson:"name" json:"name"`               //名称
	Major       []string           `bson:"major" json:"major"`             //专业
	TotalNumber int                `bson:"totalNumber" json:"totalNumber"` //考试在线人数
}
