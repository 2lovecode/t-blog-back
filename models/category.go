package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const CategoryStateNormal = 1
const CategoryStateBanned = 0

type Category struct {
	Name 		string		`json:"name" bson:"name"`
	State 		int8 		`json:"state" bson:"state"`
	AddTime 	time.Time 	`json:"addTime" bson:"addTime"`
	ModifyTime 	time.Time  	`json:"modifyTime" bson:"modifyTime"`
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