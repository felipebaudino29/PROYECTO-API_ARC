package main


import (
"log"
"os"


"github.com/gin-gonic/gin"
)


func main() {
port := os.Getenv("PORT")
if port == "" { port = "8080" }


r := gin.Default()
r.GET("/health", func(c *gin.Context) {
c.JSON(200, gin.H{"status": "ok", "service": "properties-api"})
})


log.Printf("properties-api listening on :%s", port)
if err := r.Run(":" + port); err != nil {
log.Fatal(err)
}
}