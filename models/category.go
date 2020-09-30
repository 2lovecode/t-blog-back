package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Category struct {
	Name string		`json:"name"`
	AddTime int64 	`json:"add_time"`
	ModifyTime int64  `json:"modify_time"`
}

func (cg *Category) Collection() string {
	return "category"
}

func (cg *Category)AddCategory() (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(cg.Collection()).InsertOne(context.TODO(), cg)
}

func (cg *Category)FindUserByName(name string) (user User, e error) {
	filter := bson.D{{
		"name", name,
	}}
	e = GetDb().Collection(cg.Collection()).FindOne(context.TODO(), filter).Decode(&user)
	return
}
