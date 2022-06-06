package main

import (
	"apiGateway/router"
	"log"
)

func main() {
	r := router.NewRouter()
	if err := r.Run(":8081"); err != nil {
		log.Print(err.Error())
	}
}
