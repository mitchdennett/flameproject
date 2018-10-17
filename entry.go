package main

import (
    "github.com/flame/server"
    "github.com/joho/godotenv"
)

//go:generate go run gen.go

func main() {
    err := godotenv.Load()
    if err != nil {
      
    }
    server.ListenAndServe(":8070")
}