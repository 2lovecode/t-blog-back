package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login 登录记录
type Login struct {
	AuthorID  string    `json:"-" bson:"authorID"`
	AuthCode  string    `json:"authCode" bson:"authCode"`
	Expire    int       `json:"expire" bson:"expire"`
	StartTime time.Time `json:"startTime" bson:"startTime"`
	User      User      `json:"user" bson:"user"`
}

// Collection 用户collection
func (l *Login) Collection() string {
	return "tank_login"
}

// FindUserByAuthCode 通过授权码获取用户信息
func (l *Login) FindUserByAuthCode(ctx context.Context, authCode string) (login *Login, err error) {
	// 先看有没有过期
	user := &User{}
	pipeline := mongo.Pipeline{
		bson.D{
			bson.E{
				Key: "$lookup",
				Value: bson.D{
					bson.E{Key: "from", Value: user.Collection()},
					bson.E{Key: "localField", Value: "authorID"},
					bson.E{Key: "foreignField", Value: "authorID"},
					bson.E{Key: "as", Value: "user"},
				},
			},
		},
	}

	cursor, err := GetDb().Collection(l.Collection()).Aggregate(ctx, pipeline)

	if cursor != nil {
		defer cursor.Close(ctx)
		cursor.Decode(login)
	}
	return
}

// FindLoginByAuthorID 查找
func (l *Login) FindLoginByAuthorID(ctx context.Context, authorID string) (login *Login, err error) {
	filter := bson.D{bson.E{
		Key:   "authorID",
		Value: authorID,
	}}
	cursor, err := GetDb().Collection(l.Collection()).Find(ctx, filter)
	if cursor != nil {
		defer cursor.Close(ctx)
		cursor.Decode(login)
	}
	return
}

// AddLogin 添加
func (l *Login) AddLogin(ctx context.Context) (result *mongo.InsertOneResult, err error) {
	return GetDb().Collection(l.Collection()).InsertOne(ctx, l)
}

// UpdateLoginAuthCode 更新授权码
func (l *Login) UpdateLoginAuthCode(ctx context.Context, authorCode string) (result *mongo.UpdateResult, err error) {
	login := &Login{}
	filter := bson.D{bson.E{
		Key:   "authorID",
		Value: l.AuthorID,
	}}
	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "authCode",
					Value: authorCode,
				},
				bson.E{
					Key:   "startTime",
					Value: time.Now(),
				},
			},
		},
	}
	result, err = GetDb().Collection(login.Collection()).UpdateOne(ctx, filter, update)
	return
}
