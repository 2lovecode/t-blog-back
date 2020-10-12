package models

import "time"

type ArticleList []ArticleEntry

type ArticleEntry struct {
	ArticleID 		int			`json:"articleID" bson:"articleID"`
	CID 			string 		`json:"cid" bson:"cid"`
	TID 			string 		`json:"tid" bson:"tid"`
	Title			string 		`json:"title" bson:"title"`
	Author			string		`json:"author" bson:"author"`
	AuthorAvatar 	string		`json:"authorAvatar" bson:"authorAvatar"`
	Image 			string		`json:"image" bson:"image"`
	Summary 		string		`json:"summary" bson:"summary"`
	Visited 		string		`json:"visited" bson:"visited"`
	Tags 			[]string 	`json:"tags" bson:"tags"`
	State 			int8 		`json:"state" bson:"state"`
	AddTime 		time.Time 	`json:"addTime" bson:"addTime"`
	ModifyTime 		time.Time  	`json:"modifyTime" bson:"modifyTime"`
}

func GetArticleList() ArticleList {
	articleList := ArticleList{
		ArticleEntry{ArticleID:1,},
		ArticleEntry{ArticleID:2,},
	}
	return articleList
}
