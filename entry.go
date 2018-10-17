package main

import (
	"github.com/joho/godotenv"
	"github.com/mitchdennett/flameframework/server"
)

//go:generate go run gen.go

func main() {
	err := godotenv.Load()
	if err != nil {

	}
	server.ListenAndServe(":8070")
}
