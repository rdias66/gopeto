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
	userChoice := ""
	userInput  := ""
	
	if userChoice == "data to Json" {
		requestResponse := postRequest(dataToJson(userInput))
	}
	if userChoice == "data to Excel" {
		requestResponse := postRequest(dataToExcel(userInput))
	}
}

func dataToJson (input string) string {
	return preMadeQuestion := "Transform the following string data into a organized json objects array with the format being [ { contentName: content, ...} , ...}], identify the content name via the content itself, content: " + input 
}

func dataToExcelFormat (input string) string {
	return preMadeQuestion := "Transform the following string data into a organized excel format text(to be transformed within excel to sheets), you identify what the first columns will be  by the type of content that the data provides,  content: " + input 
}

func postRequest (promptInput string) string {
	prompt := promptInput
	conversation := []string{prompt}

	// Prepare the API request payload
	data := fmt.Sprintf(`{
		"prompt": "%s",
		"max_tokens": 50,
		"temperature": 0.7,
		"n": 1,
		"stop": "\n"
	}`, strings.Join(conversation, "\n"))

	// Make the API request
	apiKey := os.Getenv("YOUR_API_KEY")
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"
	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()

	// Parse the API response
	responseBody, _ := ioutil.ReadAll(res.Body)
	return responseBody
}

