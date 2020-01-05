package main

import (
	"lms/server/database"
	"lms/server/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.GetConnection()
	router := gin.Default()

	// Setup User routes
	user.InitCtrl(db, router)

	router.Run()
}
