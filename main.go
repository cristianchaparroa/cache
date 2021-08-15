package main

import (
	"cache/app/delivery"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	s := delivery.NewServer()
	s.Run()
}
