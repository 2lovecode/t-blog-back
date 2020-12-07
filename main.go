package main

import (
	"t-blog-back/models"
	"t-blog-back/pkg/storage"
	"t-blog-back/server/http"
)

func init() {
	models.SetUp()
	storage.SetUp()
}

func main() {
	http.Start()
}
