package main

import (
	"github.com/openai/openai-go/v2"
	"github.com/joho/godotenv/cmd/godotenv"
)

// go get github.com/joho/godotenv/cmd/godotenv
// go get github.com/openai/openai-go/v2
// go get github.com/google/uuid


func main() {
	err := godotenv.Load()
  	if err != nil {
  	 	log.Fatal("Error loading .env file")
  	}
	
	openaiClient := openai.NewClient(os.Getenv("YOUR_API_KEY"))
	// ...
}
