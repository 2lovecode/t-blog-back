package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/setting"
	"t-blog-back/routers"
)

var Db *gorm.DB

func main() {
	router := routers.InitRouter()

	Db, dbErr := gorm.Open(setting.DbCfg.Type, setting.DbCfg.Name)

	if dbErr != nil {
		log.Fatalf("start error: %v", dbErr)
	}
	defer Db.Close()

	models.InitTankDb(Db)

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