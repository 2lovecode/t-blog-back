package models

import "time"

// ArticleList 文章列表
type ArticleList []ArticleEntry

// ArticleEntry 文章
type ArticleEntry struct {
	ArticleID    int       `json:"articleID" bson:"articleID"`
	CategoryID   string    `json:"category_id" bson:"category_id"`
	TagID        string    `json:"tag_id" bson:"category_id"`
	AuthorID     string    `json:"author_id" bson:"author_id"`
	Title        string    `json:"title" bson:"title"`
	Author       string    `json:"author" bson:"author"`
	AuthorAvatar string    `json:"authorAvatar" bson:"authorAvatar"`
	Image        string    `json:"image" bson:"image"`
	Summary      string    `json:"summary" bson:"summary"`
	Visited      string    `json:"visited" bson:"visited"`
	Tags         []string  `json:"tags" bson:"tags"`
	State        int8      `json:"state" bson:"state"`
	AddTime      time.Time `json:"addTime" bson:"addTime"`
	ModifyTime   time.Time `json:"modifyTime" bson:"modifyTime"`
}

// GetArticleList 文章列表
func GetArticleList() ArticleList {
	articleList := ArticleList{
		ArticleEntry{ArticleID: 1},
		ArticleEntry{ArticleID: 2},
	}
	return articleList
}
