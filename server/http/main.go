package http

import (
	"fmt"
	"log"
	"net/http"
	"t-blog-back/pkg/setting"
	"t-blog-back/routers"
)

func Start() {
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
		log.Fatalf("http listen err: %v", err)
	}
}
