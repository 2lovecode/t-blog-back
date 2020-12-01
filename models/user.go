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
	AuthorID   string    `json:"authorID" bson:"authorID"`
	Name       string    `json:"name" bson:"name"`
	Pass       string    `json:"pass" bson:"pass"`
	NickName   string    `json:"nickName" bson:"nickName"`
	Avatar     string    `json:"avatar" bson:"avatar"`
	State      int8      `json:"state" bson:"state"`
	AddTime    time.Time `json:"addTime" bson:"addTime"`
	ModifyTime time.Time `json:"modifyTime" bson:"modifyTime"`
}

// Collection 用户collection
func (u *User) Collection() string {
	return "tank_user"
}

// AddUser 添加用户
func (u *User) AddUser(ctx context.Context) (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(u.Collection()).InsertOne(ctx, u)
}

// UpdateUser 添加用户
func (u *User) UpdateUser(ctx context.Context) (result *mongo.UpdateResult, e error) {
	filter := bson.D{primitive.E{
		Key:   "name",
		Value: u.Name,
	}}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "authorID",
					Value: u.AuthorID,
				},
				primitive.E{
					Key:   "pass",
					Value: u.Pass,
				},
				primitive.E{
					Key:   "modifyTime",
					Value: u.ModifyTime,
				},
			},
		},
	}
	return GetDb().Collection(u.Collection()).UpdateOne(ctx, filter, update)
}

// FindUserByName 用户名查询用户
func (u *User) FindUserByName(ctx context.Context, name string) (e error) {
	filter := bson.D{primitive.E{
		Key:   "name",
		Value: name,
	}}
	e = GetDb().Collection(u.Collection()).FindOne(ctx, filter).Decode(&u)
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
