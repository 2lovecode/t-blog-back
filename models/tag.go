package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const TagStateNormal = 1
const TagStateBanned = 0

type Tag struct {
	ID 		string 	`json:"id" bson:"id"`
	Name 	string 	`json:"name" bson:"name"`
	State 	int 	`json:"state" bson:"state"`
	AddTime 	time.Time 	`json:"addTime" bson:"addTime"`
	ModifyTime 	time.Time  	`json:"modifyTime" bson:"modifyTime"`
}

func (tg *Tag) Collection() string {
	return "tag"
}

func (tg *Tag)AddTag() (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(tg.Collection()).InsertOne(context.TODO(), tg)
}

func (tg *Tag)FindByID(id string) (tag Tag, e error) {
	filter := bson.D{{
		"id", id,
	}}
	e = GetDb().Collection(tg.Collection()).FindOne(context.TODO(), filter).Decode(&tag)
	return
}

func (tg *Tag)FindByName(name string) (tag Tag, e error) {
	filter := bson.D{{
		"name", name,
	}}
	e = GetDb().Collection(tg.Collection()).FindOne(context.TODO(), filter).Decode(&tag)
	return
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tag Tag) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := tankDb.Collection("tag")
	res, _ :=  collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	fmt.Println(res)
	tag = Tag{State:12}
	return
}


