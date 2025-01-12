package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func getLogIn() gin.Accounts {

	password := os.Getenv("ADMIN_PASSWORD")

	if password == "" {
		fmt.Println("ADMIN_PASSWORD not set")
		log.Fatal("ADMIN_PASSWORD not set")
	}

	return gin.Accounts{
		"Admin": password,
	}
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "JOS - JSON Object Store\n")
}

func uploadFile(c *gin.Context) {
	// Get the file from the post context
	file, err := c.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
	}

	// Open the file
	openFile, err := file.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer openFile.Close()

	// Read the contents of the file
	fileContents, err := io.ReadAll(openFile)
	if err != nil {
		log.Println("Error reading file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// Generate hash of the file contents
	file_hash := sha256.Sum256([]byte(fileContents))
	file_hash_filename := string(file_hash[:])

	// Save the file to the out directory
	c.SaveUploadedFile(file, "out/" + file_hash_filename)

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded and saved successfully!",
		"filename": file_hash_filename,
	})
}

func uploadJSON(c *gin.Context) {
	// Dynamic JSON structure
	var jsonData map[string]interface{}

	// Bind incoming JSON to the map
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert JSON to a byte slice
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize JSON"})
		return
	}

	json_hash := sha256.Sum256([]byte(jsonBytes))
	json_hash_filename := string(json_hash[:])

	// Write JSON file out
	err = os.WriteFile("out/" + json_hash_filename, jsonBytes, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save JSON to file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "JSON uploaded and saved successfully!",
		"filename": json_hash_filename,
	})
}

func main() {
	router := gin.Default()

	routeStart := "/api/v1/"
	authAccount := getLogIn()

	authedSubRoute := router.Group(routeStart, gin.BasicAuth(authAccount))

	authedSubRoute.GET("/", homePage)
	authedSubRoute.POST("/upload-file", uploadFile)
	authedSubRoute.POST("/upload-json", uploadJSON)

	router.Run(":8080")
}
