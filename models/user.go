package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// UpdateUser 更改用户
func (u *User) UpdateUser(ctx context.Context) (result *mongo.UpdateResult, e error) {
	filter := bson.D{
		bson.E{
			Key:   "name",
			Value: u.Name,
		},
	}
	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "authorID",
					Value: u.AuthorID,
				},
				bson.E{
					Key:   "pass",
					Value: u.Pass,
				},
				bson.E{
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
	filter := bson.D{
		bson.E{
			Key:   "name",
			Value: name,
		},
	}
	e = GetDb().Collection(u.Collection()).FindOne(ctx, filter).Decode(&u)
	return
}

// FindUserByNameAndPass 用户名密码查询用户
func (u *User) FindUserByNameAndPass(ctx context.Context, name string, pass string) (user *User, e error) {
	filter := bson.D{
		bson.E{
			Key:   "name",
			Value: name,
		},
		bson.E{
			Key:   "pass",
			Value: pass,
		},
	}
	e = GetDb().Collection(u.Collection()).FindOne(ctx, filter).Decode(&user)
	return
}

// FindUserByAuthorID AuthorID查询用户
func (u *User) FindUserByAuthorID(ctx context.Context, authorID string) (user *User, err error) {
	filter := bson.D{
		bson.E{
			Key:   "authorID",
			Value: authorID,
		},
	}
	user = &User{}
	err = GetDb().Collection(u.Collection()).FindOne(ctx, filter).Decode(user)
	return
}
