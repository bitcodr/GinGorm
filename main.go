package main

import (
	"os"

	"github.com/amiralii/gin-gorm/config"
	"github.com/amiralii/gin-gorm/aggregates/note"
	"github.com/gin-gonic/gin"
)

func main() {

	gin := gin.Default()

	config.Environment()

	config.InitDB()
	
	note.Routes(gin)

	gin.Run(os.Getenv("APP_PORT"))
}