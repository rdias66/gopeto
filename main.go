package main

import (
	"github.com/openai/openai-go/v2"
	"github.com/joho/godotenv/cmd/godotenv"
)

// must run on first download
// go mod init 
// go get github.com/joho/godotenv/cmd/godotenv
// go get github.com/openai/openai-go/v2

func main() {
	err := godotenv.Load()
  	if err != nil {
  	 	log.Fatal("Error loading .env file")
  	}	
}

func dataToJson (input string) string {
	return preMadeQuestion := "Transform the following string data into a organized json objects array with the format being [ { contentName: content, ...} , ...}], identify the content name via the content itself, content: " + input 
}

func dataToExcelFormat (input string) string {
	return preMadeQuestion := "Transform the following string data into a organized excel format text(to be transformed within excel to sheets), you identify what the first columns will be  by the type of content that the data provides,  content: " + input 
}


