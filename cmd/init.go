package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"t-blog-back/models"
	"time"
)

func init() {
	models.SetUp()
}

func main() {
	now := time.Now().Unix()

	user := models.User{}
	name := "admin"
	if u, e := user.FindUserByName(name); e != nil {
		if pass, e := bcrypt.GenerateFromPassword([]byte("admin"), 16); e == nil {
			user.Name = name
			user.Pass = string(pass)
			user.AddTime = now
			user.ModifyTime = now

			r, e := user.AddUser()
			fmt.Println("result:", r)
			fmt.Println("error:", e)
		}
	} else {
		fmt.Println(u, e)
	}
}
