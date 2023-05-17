package main

import (
	"github.com/openai/openai-go/v2"
)

// go get go get github.com/joho/godotenv/cmd/godotenv
// go get github.com/openai/openai-go/v2



func main() {
	openaiClient := openai.NewClient("YOUR_API_KEY")
	// ...
}
