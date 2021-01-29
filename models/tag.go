package models

import (
	"context"
	"log"
	"t-blog-back/pkg/utils"
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
	return "tank_tag"
}

// AddTag 添加标签
func (tg *Tag) AddTag(ctx context.Context) (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(tg.Collection()).InsertOne(ctx, tg)
}

// FindByID id查询
func (tg *Tag) FindByID(ctx context.Context, id string) (tag Tag, e error) {
	filter := bson.D{primitive.E{
		Key:   "id",
		Value: id,
	}}
	e = GetDb().Collection(tg.Collection()).FindOne(ctx, filter).Decode(&tag)
	return
}

// FindByName name查询
func (tg *Tag) FindByName(ctx context.Context, name string) (tag Tag, e error) {
	filter := bson.D{primitive.E{
		Key:   "name",
		Value: name,
	}}
	e = GetDb().Collection(tg.Collection()).FindOne(ctx, filter).Decode(&tag)
	return
}

// GetTags 获取标签
func (tg *Tag)GetTags(ctx context.Context, page *utils.Pagination, filter bson.D) (tags []Tag, err error) {
	collection :=  GetDb().Collection(tg.Collection())
	cnt, err := collection.CountDocuments(ctx, filter)

	if err == nil && cnt > 0 {
		findOptions := page.GetFindOptions()
		cursor, err := collection.Find(ctx, filter, findOptions)
		if err == nil {
			for cursor.Next(ctx) {
				// 创建一个值，将单个文档解码为该值
				var elem Tag
				e := cursor.Decode(&elem)
				if e == nil {
					tags = append(tags, elem)
				}
			}

			if err = cursor.Err(); err != nil {
				log.Fatal(err)
			}
		}
		page.SetTotal(cnt)
	}
	return
}
