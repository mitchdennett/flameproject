package main

import (
    "github.com/flame/server"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
      
    }
    server.ListenAndServe(":8070")
}