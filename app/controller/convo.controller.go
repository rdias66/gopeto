package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go/v2"
	"github.com/google/uuid"
	//"github.com/rdias/gopeto/database" to be made
	"github.com/rdias66/gopeto/app/models/models.schema.go"
)

type ConvoController struct {
	OpenaiClient *openai.Client
}

func (ConvoController *ConvoController) CreateConvo(Context *gin.Context) {
	// Retrieve user input from the request
	userInput := Context.PostForm("user_input")

	// Call the OpenAI API to generate a model response
	resp, err := ConvoController.OpenaiClient.Completions.Create(
		&openai.CompletionCreateRequest{
			Model:     "gpt-3.5-turbo",
			MaxTokens: 50,
			Prompt:    userInput,
		},
	)
	if err != nil {
		Context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate model response"})
		return
	}

	// Retrieve the generated response from the API response
	modelResponse := resp.Choices[0].Text

	// Store the conversation in the database
	convo := models.Convo{
		Id:   	     uuid.New(),
		UserInput:   userInput,
		ModelOutput: modelResponse,
		Timestamp:   time.Now(),
	}
	//store it here. to be made

	// Return the model response in the API response
	Context.JSON(http.StatusOK, gin.H{"model_response": modelResponse})
}
