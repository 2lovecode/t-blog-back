package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"t-blog-back/pkg/setting"
	"time"
)

type Model struct {
}

var tankDbClient *mongo.Client
var tankDb  *mongo.Database

func SetUp() {
	dbCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbClient, dbErr := mongo.Connect(dbCtx, options.Client().SetAuth(options.Credential{
		Username:                setting.DbCfg.User,
		Password:                setting.DbCfg.Pass,
	}).ApplyURI(setting.DbCfg.Host))

	if dbErr != nil {
		log.Fatalf("start error: %v", dbErr)
	}
	dbErr = dbClient.Ping(dbCtx, readpref.Primary())
	if dbErr != nil {
		log.Fatalf("db connection  error: %v", dbErr)
	}

	log.Println("connect db success!")

	tankDbClient = dbClient
	tankDb = tankDbClient.Database(setting.DbCfg.Name)
}

func GetDb() *mongo.Database {
	return tankDb
}

func GetClient() *mongo.Client {
	return tankDbClient
}
