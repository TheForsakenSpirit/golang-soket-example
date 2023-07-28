package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	server "example/socket/Server"
)

func main() {
	fmt.Println("Hello World")

	engine := gin.Default()
	server.BuildRouter(engine)
	engine.Run(":8080")
}
