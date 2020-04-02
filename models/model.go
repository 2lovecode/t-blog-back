package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"t-blog-back/pkg/setting"
)

type Model struct {
}

var tankDbClient *mongo.Client
var tankDb  *mongo.Database

func InitTankDb(dbClient *mongo.Client) {
	tankDbClient = dbClient
	tankDb = tankDbClient.Database(setting.DbCfg.Name)
}
