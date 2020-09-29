package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name string		`json:"name"`
	Pass string		`json:"pass"`
	Token string 	`json:"token"`
	Avatar string 	`json:"avatar"`
	AddTime int64 	`json:"add_time"`
	ModifyTime int64  `json:"modify_time"`
}

func (u *User) Collection() string {
	return "user"
}

func (u *User)AddUser() (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(u.Collection()).InsertOne(context.TODO(), u)
}

func (u *User)FindUserByName(name string) (user User, e error) {
	filter := bson.D{{
		"name", name,
	}}
	e = GetDb().Collection(u.Collection()).FindOne(context.TODO(), filter).Decode(&user)
	return
}

func (u *User)UpdateToken(name string, token string) (result *mongo.UpdateResult, e error) {
	filter := bson.D{{
		"name", name,
	}}
	update := bson.D{
		{"$set", bson.D{{"token", token}}},
	}
	result, e = GetDb().Collection(u.Collection()).UpdateOne(context.TODO(), filter, update)
	return
}

func (u *User)FindUserByNameAndPass(name string, pass string) (user User, e error) {
	filter := bson.D{{
		"name", name,
	},{
		"pass", pass,
	}}
	e = GetDb().Collection(u.Collection()).FindOne(context.TODO(), filter).Decode(&user)
	return
}

func (u *User)FindUserByToken(token string) (user User, e error) {
	filter := bson.D{{
		"token", token,
	}}
	e = GetDb().Collection(u.Collection()).FindOne(context.TODO(), filter).Decode(&user)
	return
}