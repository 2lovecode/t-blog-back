package models


type ArticleList []ArticleEntry

type ArticleEntry struct {
	ArticleID 		int			`json:"articleID"`
	Title			string 		`json:"title"`
	Author			string		`json:"author"`
	AuthorAvatar 	string		`json:"authorAvatar"`
	CreatedTime 	string		`json:"createdTime"`
	Image 			string		`json:"image"`
	Summary 		string		`json:"summary"`
	Visited 		string		`json:"visited"`
	SkipUrl 		string		`json:"skipUrl"`
	Tags 			[]string 	`json:"tags"`
}

type Article struct {
	Model
}

func GetArticleList() ArticleList {
	articleList := ArticleList{
		ArticleEntry{ArticleID:1,},
		ArticleEntry{ArticleID:2,},
	}
	return articleList
}
