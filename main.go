package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/setting"
	"t-blog-back/routers"
	"time"
)


func main() {
	router := routers.InitRouter()

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

	models.InitTankDb(dbClient)

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()

	if err != nil {
		log.Fatalf("start error: %v", err)
	}

}