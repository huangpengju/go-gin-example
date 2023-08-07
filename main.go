package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "helleo world!",
		})
	})
	r.Run()
	// https://eddycjy.com/posts/go/gin/2018-02-11-api-01/
}
