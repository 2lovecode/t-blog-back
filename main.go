package main

import (
	"fmt"
	"log"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/setting"
	"t-blog-back/routers"
)

func init() {
	models.SetUp()
}

func main() {
	router := routers.InitRouter()

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