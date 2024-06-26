package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User 模型
type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Account      string               `bson:"account"    json:"account"    form:"account"`
	Password     string               `bson:"password"   json:"password"   form:"password"`
	Name         string               `bson:"name"       json:"name"       form:"name"`
	AvatarURL    string               `bson:"avatar_url" json:"avatar_url" form:"avatar_url" `
	Sex          string               `bson:"sex"        json:"sex"        form:"sex"           validate:"Sex"`
	Phone        string               `bson:"phone"      json:"phone"      form:"phone" `
	IsValid      string               `bson:"is_valid"   json:"is_valid"   form:"is_valid" `
	Salt         string               `bson:"salt"       json:"salt"       form:"salt" `
	Info         string               `bson:"info" json:"info"`                                 //考试在线状态：0为离线 1为在线
	Major        string               `bson:"major" json:"major"`                                   //专业
	Authority    primitive.ObjectID   `bson:"authority" json:"authority" form:"authority" `         //存数据库的角色
	Verification []primitive.ObjectID `bson:"verification" json:"verification" form:"verification"` //访问数据的权限
	Role         string               `bson:"role" json:"role" form:"role" validate:"StringNumber"` //返回给前端的用户角色---与Authority相对应
}

// Login 模型
type Login struct {
	Account  string `  bson:"account"  form:"account" json:"account"  `
	Password string `  bson:"password" form:"password" json:"password" `
}
