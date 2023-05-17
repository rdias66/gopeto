package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go/v2"

	"github.com/yourusername/projectname/database"
	"github.com/yourusername/projectname/models"
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
		UserInput:   userInput,
		ModelOutput: modelResponse,
		Timestamp:   time.Now(),
	}
	if err := database.CreateConvo(&convo); err != nil {
		Context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create conversation"})
		return
	}

	// Return the model response in the API response
	Context.JSON(http.StatusOK, gin.H{"model_response": modelResponse})
}
