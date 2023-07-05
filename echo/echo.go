package main

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func recoverPanic() {
    if err := recover(); err != nil {
        fmt.Printf("RECOVERED: %v\n", err)
    }
}