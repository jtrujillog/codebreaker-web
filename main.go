package main

import (
  "net/http"
  "os"
  "github.com/gin-gonic/gin"
)

func main(){
  port := os.Getenv("PORT")

  if port == "" {
    port = "8081"
  }

  router := gin.Default()

  router.GET("/codebreaker/setup/:number", func(c *gin.Context){
    number := c.Param("number")
    setCode(number)
    c.String(http.StatusOK, "Secret number configured: " + number)
  })

  router.GET("/codebreaker/guess/:number", func(c *gin.Context){
    number := c.Param("number")
    result := puntoFama(number)
    c.String(http.StatusOK, "Answer: " + result)
  })

  router.Run(":" +port)

}
