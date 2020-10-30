package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User 用户
type User struct {
	Name       string    `json:"name" bson:"name"`
	Pass       string    `json:"pass" bson:"pass"`
	Token      string    `json:"token" bson:"token"`
	Avatar     string    `json:"avatar" bson:"avatar"`
	State      int8      `json:"state" bson:"state"`
	AddTime    time.Time `json:"addTime" bson:"addTime"`
	ModifyTime time.Time `json:"modifyTime" bson:"modifyTime"`
}

// Collection 用户collection
func (u *User) Collection() string {
	return "user"
}

// AddUser 添加用户
func (u *User) AddUser() (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(u.Collection()).InsertOne(context.TODO(), u)
}

// FindUserByName 用户名查询用户
func (u *User) FindUserByName(name string) (user User, e error) {
	filter := bson.D{primitive.E{
		Key:   "name",
		Value: name,
	}}
	e = GetDb().Collection(u.Collection()).FindOne(context.TODO(), filter).Decode(&user)
	return
}

// UpdateToken 更新token
func (u *User) UpdateToken(name string, token string) (result *mongo.UpdateResult, e error) {
	filter := bson.D{primitive.E{
		Key:   "name",
		Value: name,
	}}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{primitive.E{
				Key:   "token",
				Value: token,
			},
			},
		},
	}
	result, e = GetDb().Collection(u.Collection()).UpdateOne(context.TODO(), filter, update)
	return
}

// FindUserByNameAndPass 用户名密码查询用户
func (u *User) FindUserByNameAndPass(name string, pass string) (user User, e error) {
	filter := bson.D{primitive.E{
		Key:   "name",
		Value: name,
	}, primitive.E{
		Key:   "pass",
		Value: pass,
	}}
	e = GetDb().Collection(u.Collection()).FindOne(context.TODO(), filter).Decode(&user)
	return
}

// FindUserByToken token查询用户
func (u *User) FindUserByToken(token string) (user User, e error) {
	filter := bson.D{primitive.E{
		Key:   "token",
		Value: token,
	}}
	e = GetDb().Collection(u.Collection()).FindOne(context.TODO(), filter).Decode(&user)
	return
}
