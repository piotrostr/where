package main

import (
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/profiler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	return r
}

func SetupProfiler() {
	cfg := profiler.Config{
		Service:        "hello-world",
		ServiceVersion: "1.0.0",
		ProjectID:      "piotrostr-resources",
	}
	if err := profiler.Start(cfg); err != nil {
		errMsg := fmt.Sprintf("failed to start profiler: %v", err)
		log.Fatal(errMsg)
	}
}

func main() {
	port := *flag.String("port", "8080", "port to run on")
	flag.Parse()

	SetupProfiler()

	r := SetupRouter()
	log.Printf("Listening on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
