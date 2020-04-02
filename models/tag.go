package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Tag struct {
	Model
	Name 	string 	`json:"name"`
	State 	int 	`json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tag Tag) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := tankDb.Collection("tag")
	res, _ :=  collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	fmt.Println(res)
	tag = Tag{State:12}
	return
}

func GetTagTotal(maps interface{}) (count int) {
	return
}


