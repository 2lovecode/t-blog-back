package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	Desc       string    `json:"desc" bson:"desc"`
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
	filter := bson.D{
		bson.E{
			Key:   "id",
			Value: id,
		},
	}
	e = GetDb().Collection(cg.Collection()).FindOne(context.TODO(), filter).Decode(&category)
	return
}

// FindByName name查询
func (cg *Category) FindByName(ctx context.Context, name string) (err error) {
	filter := bson.D{
		bson.E{
			Key:   "name",
			Value: name,
		},
	}
	err = GetDb().Collection(cg.Collection()).FindOne(ctx, filter).Decode(cg)
	return
}

// IsEmpty 数据是否为空
func (cg *Category) IsEmpty() bool {
	return cg.ID == ""
}

// FindAll 返回所有分类
func (cg *Category) FindAll(ctx context.Context) (categories []Category, err error) {
	filter := bson.D{}

	cursor, err := GetDb().Collection(cg.Collection()).Find(ctx, filter)
	if cursor != nil {
		for cursor.Next(ctx) {
			category := Category{}
			cursor.Decode(&category)
			categories = append(categories, category)
		}
	}
	return
}
