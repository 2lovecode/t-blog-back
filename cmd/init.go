package main

import (
	"context"
	"fmt"
	"t-blog-back/logic/user"
	"t-blog-back/models"
)

func init() {
	models.SetUp()
}

func main() {
	iData, err := user.InitUser(context.Background(), user.InitUserForm{})

	fmt.Println(iData, err)
}
