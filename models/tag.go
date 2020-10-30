package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TagStateNormal 正常
const TagStateNormal = 1

// TagStateBanned 禁用
const TagStateBanned = 0

// Tag 标签
type Tag struct {
	ID         string    `json:"id" bson:"id"`
	Name       string    `json:"name" bson:"name"`
	State      int       `json:"state" bson:"state"`
	AddTime    time.Time `json:"addTime" bson:"addTime"`
	ModifyTime time.Time `json:"modifyTime" bson:"modifyTime"`
}

// Collection 标签collection
func (tg *Tag) Collection() string {
	return "tag"
}

// AddTag 添加标签
func (tg *Tag) AddTag() (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(tg.Collection()).InsertOne(context.TODO(), tg)
}

// FindByID id查询
func (tg *Tag) FindByID(id string) (tag Tag, e error) {
	filter := bson.D{primitive.E{
		Key:   "id",
		Value: id,
	}}
	e = GetDb().Collection(tg.Collection()).FindOne(context.TODO(), filter).Decode(&tag)
	return
}

// FindByName name查询
func (tg *Tag) FindByName(name string) (tag Tag, e error) {
	filter := bson.D{primitive.E{
		Key:   "name",
		Value: name,
	}}
	e = GetDb().Collection(tg.Collection()).FindOne(context.TODO(), filter).Decode(&tag)
	return
}

// GetTags 获取标签
func GetTags(ctx context.Context, pageNum int, pageSize int, maps interface{}) (tag Tag) {
	collection := tankDb.Collection("tag")
	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	fmt.Println(res)
	tag = Tag{State: 12}
	return
}
