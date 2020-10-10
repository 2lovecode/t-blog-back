package main

import (
	"t-blog-back/models"
	"t-blog-back/server/http"
)

func init() {
	models.SetUp()
}

func main() {
	http.Start()
}