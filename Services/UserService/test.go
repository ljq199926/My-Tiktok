package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123456"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败:", err)
	}
	encodePWD := string(hash)
	fmt.Println(encodePWD)
}
