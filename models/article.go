package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Article 文章
type Article struct {
	ArticleID    string    `json:"articleID" bson:"articleID"`
	AuthorID     string    `json:"authorID" bson:"authorID"`
	Title        string    `json:"title" bson:"title"`
	Author       string    `json:"author" bson:"author"`
	AuthorAvatar string    `json:"authorAvatar" bson:"authorAvatar"`
	Image        string    `json:"image" bson:"image"`
	Summary      string    `json:"summary" bson:"summary"`
	Content		 string    `json:"content" bson:"content"`
	Visited      string    `json:"visited" bson:"visited"`
	Tags         []string  `json:"tags" bson:"tags"`
	State        int8      `json:"state" bson:"state"`
	AddTime      time.Time `json:"addTime" bson:"addTime"`
	ModifyTime   time.Time `json:"modifyTime" bson:"modifyTime"`
}

// Collection 标签collection
func (at *Article) Collection() string {
	return "tank_article"
}

// GetArticleList 文章列表
func (at *Article) GetArticles(ctx context.Context, pageNum int, pageSize int, maps map[string]interface{}) (articles []Article) {
	filter := bson.D{primitive.E{}}
	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))

	collection :=  GetDb().Collection(at.Collection())
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err == nil {
		for cursor.Next(ctx) {
			// 创建一个值，将单个文档解码为该值
			var elem Article
			e := cursor.Decode(&elem)
			if e == nil {
				articles = append(articles, elem)
			}

		}

		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return
}

// AddArticle 添加文章
func (at *Article) AddArticle(ctx context.Context) (result *mongo.InsertOneResult, e error) {
	return GetDb().Collection(at.Collection()).InsertOne(ctx, at)
}


