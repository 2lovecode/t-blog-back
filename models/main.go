package models

import (
	"context"
	"log"
	"t-blog-back/pkg/setting"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Model 模型
type Model struct {
}

var tankCancelFunc context.CancelFunc
var tankDbClient *mongo.Client
var tankDb *mongo.Database

// SetUp mongodb初始化
func SetUp() {
	dbCtx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	dbClient, dbErr := mongo.Connect(dbCtx, options.Client().SetAuth(options.Credential{
		Username: setting.DbCfg.User,
		Password: setting.DbCfg.Pass,
	}).SetConnectTimeout(3*time.Second).ApplyURI(setting.DbCfg.Host))

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
	tankCancelFunc = cancelFunc
}

// GetDb database
func GetDb() *mongo.Database {
	return tankDb
}

// GetClient client
func GetClient() *mongo.Client {
	return tankDbClient
}
