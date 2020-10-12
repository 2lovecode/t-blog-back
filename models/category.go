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
	ID 			string 		`json:"id" bson:"id"`
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

func (cg *Category)FindByID(id string) (category Category, e error) {
	filter := bson.D{{
		"id", id,
	}}
	e = GetDb().Collection(cg.Collection()).FindOne(context.TODO(), filter).Decode(&category)
	return
}

func (cg *Category)FindByName(name string) (category Category, e error) {
	filter := bson.D{{
		"name", name,
	}}
	e = GetDb().Collection(cg.Collection()).FindOne(context.TODO(), filter).Decode(&category)
	return
}
