package models


type Tag struct {
	Model
	Name 	string 	`json:"name"`
	State 	int 	`json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tag Tag) {
	TankDb.First(&tag, 1)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	return
}


