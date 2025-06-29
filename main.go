// main.go

package main

import (
	"github.com/gin-gonic/gin"

	"github.com/thamruicong/Gin-Golang/handlers"
)

var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("templates/*")

   // Handle Index
   router.GET("/", handlers.ShowIndexPage)

  // Start serving the application
  router.Run()

}