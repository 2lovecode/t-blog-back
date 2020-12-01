package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login 登录记录
type Login struct {
	AuthorID string    `json:"-" bson:"authorID"`
	Token    string    `json:"token" bson:"token"`
	Expire   int       `json:"expire" bson:"expire"`
	AddTime  time.Time `json:"addTime" bson:"addTime"`
	User     User      `json:"user" bson:"user"`
}

// Collection 用户collection
func (l *Login) Collection() string {
	return "tank_login"
}

// FindUserByToken 通过token获取用户信息
func (l *Login) FindUserByToken(ctx context.Context, token string) (user *User, err error) {
	user = &User{}
	login := Login{}
	pipeline := mongo.Pipeline{
		bson.D{
			primitive.E{
				Key: "$lookup",
				Value: bson.D{
					primitive.E{Key: "from", Value: user.Collection()},
					primitive.E{Key: "localField", Value: "authorID"},
					primitive.E{Key: "foreignField", Value: "authorID"},
					primitive.E{Key: "as", Value: "user"},
				},
			},
		},
	}

	cursor, err := GetDb().Collection(l.Collection()).Aggregate(ctx, pipeline)

	cursor.Decode(&login)

	defer cursor.Close(ctx)

	return
}

// FindLoginByAuthorID 查找
func (l *Login) FindLoginByAuthorID(ctx context.Context, authorID string) (login *Login, err error) {
	filter := bson.D{primitive.E{
		Key:   "authorID",
		Value: l.AuthorID,
	}}
	cursor, err := GetDb().Collection(l.Collection()).Find(ctx, filter)
	cursor.Decode(login)
	return
}

// UpdateToken 更新token
func (l *Login) CreateLogin(ctx context.Context, token string) (result *mongo.UpdateResult, err error) {
	login := &Login{}
	filter := bson.D{primitive.E{
		Key:   "authorID",
		Value: l.AuthorID,
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
	_, err = GetDb().Collection(login.Collection()).UpdateOne(ctx, filter, update)
	return
}

// UpdateToken 更新token
func (l *Login) UpdateToken(ctx context.Context, token string) (result *mongo.UpdateResult, err error) {
	login := &Login{}
	filter := bson.D{primitive.E{
		Key:   "authorID",
		Value: l.AuthorID,
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
	_, err = GetDb().Collection(login.Collection()).UpdateOne(ctx, filter, update)
	return
}
