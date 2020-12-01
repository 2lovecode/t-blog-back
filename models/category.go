package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CategoryStateNormal 正常
const CategoryStateNormal = 1

// CategoryStateBanned 禁用
const CategoryStateBanned = 0

// Category 分类
type Category struct {
	ID         string    `json:"id" bson:"id"`
	AuthorID   string    `json:"-" bson:"authorID"`
	Name       string    `json:"name" bson:"name"`
	State      int8      `json:"state" bson:"state"`
	AddTime    time.Time `json:"addTime" bson:"addTime"`
	ModifyTime time.Time `json:"modifyTime" bson:"modifyTime"`
}

// Collection 分类collection
func (cg *Category) Collection() string {
	return "tank_category"
}

// AddCategory 添加
func (cg *Category) AddCategory(ctx context.Context) (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(cg.Collection()).InsertOne(ctx, cg)
}

// FindByID id查询
func (cg *Category) FindByID(id string) (category Category, e error) {
	filter := bson.D{primitive.E{
		Key:   "id",
		Value: id,
	}}
	e = GetDb().Collection(cg.Collection()).FindOne(context.TODO(), filter).Decode(&category)
	return
}

// FindByName name查询
func (cg *Category) FindByName(ctx context.Context, name string) (e error) {
	filter := bson.D{primitive.E{
		Key:   "name",
		Value: name,
	}}
	e = GetDb().Collection(cg.Collection()).FindOne(ctx, filter).Decode(&cg)
	return
}

// IsEmpty 数据是否为空
func (cg *Category) IsEmpty() bool {
	return cg.ID == ""
}
