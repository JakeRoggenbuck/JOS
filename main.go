package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"fmt"
	"net/http"
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
	c.String(http.StatusOK, "Home\n")
}

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
	}

	log.Println(file.Filename)

	c.SaveUploadedFile(file, "out/uploaded-"+file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func main() {
	router := gin.Default()

	routeStart := "/api/v1/"
	authAccount := getLogIn()

	authedSubRoute := router.Group(routeStart, gin.BasicAuth(authAccount))

	authedSubRoute.GET("/", homePage)
	authedSubRoute.POST("/upload", uploadFile)

	router.Run(":8080")
}
