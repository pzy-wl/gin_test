package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

// User contains user information
type UserInfo struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`       //必选项
	Age       uint8  `validate:"gte=0,lte=100"`  //0<=age<=100
	Email     string `validate:"required,email"` //邮箱格式， 必选项
}

func main() {
	validate := validator.New()
	user := &UserInfo{
		FirstName: "Badger",
		LastName:  "Smith",
		Age:       12,
		Email:     "123@gamail.com",
	}
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return
	}
	fmt.Printf("用户信息是:%v\n", user)
	fmt.Println("success")
}
