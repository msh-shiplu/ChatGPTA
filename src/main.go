package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a POST route for the API endpoint
	router.POST("/feedback", defaultFeedbackHandler)

	router.POST("/hints", hintsFeedbackHandler)

	router.POST("/example", exampleFeedbackHandler)

	router.POST("/explanation", explanationFeedbackHandler)

	router.POST("/encourage_with_hints", encourageWithHintsFeedbackHandler)

	router.POST("/encourage_with_example", encourageWithExampleFeedbackHandler)

	router.POST("/encourage_with_explanation", encourageWithExplanationFeedbackHandler)

	router.POST("/mixed_prompt", allTypeFeedbackHandler)

	router.POST("/instructions_with_example", effectiveFeedbackInstructionWithExample)

	router.POST("/instructions_without_example", effectiveFeedbackInstructionWithoutExample)

	// Start the server
	router.Run("localhost:8000")
}

func processHandlers(c *gin.Context, f getPromptFunc) {
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestData := createRequestData(f(requestBody.Course, requestBody.Duration), requestBody)
	makeRequest(c, requestData)
}

func defaultFeedbackHandler(c *gin.Context) {
	processHandlers(c, getDefaultPrompt)
}

func hintsFeedbackHandler(c *gin.Context) {
	processHandlers(c, getHintsPrompt)
}

func exampleFeedbackHandler(c *gin.Context) {
	processHandlers(c, getGiveExamplePrompt)
}

func explanationFeedbackHandler(c *gin.Context) {
	processHandlers(c, getExplanationPrompt)
}

func encourageWithHintsFeedbackHandler(c *gin.Context) {
	processHandlers(c, getEcouragementWithHintsPrompt)
}

func encourageWithExampleFeedbackHandler(c *gin.Context) {
	processHandlers(c, getEcouragementWithExamplePrompt)
}

func encourageWithExplanationFeedbackHandler(c *gin.Context) {
	processHandlers(c, getEcouragementWithExplanationPrompt)
}

func allTypeFeedbackHandler(c *gin.Context) {
	processHandlers(c, getAllNonDefaultPrompt)
}

func effectiveFeedbackInstructionWithExample(c *gin.Context) {
	processHandlers(c, getFeedbackWithInstructionsWithExample)
}

func effectiveFeedbackInstructionWithoutExample(c *gin.Context) {
	processHandlers(c, getFeedbackWithInstructionsWithoutExample)
}
