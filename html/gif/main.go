package main

import (
	"github.com/maxzerbini/packagemain/html/gif/animation"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.New()
    // Global middleware
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
	router.GET("/animated.gif", animatedgif)
	router.Run("0.0.0.0:8000")
}

func animatedgif (c *gin.Context) {
	name := c.Query("name")
	animation.GenerateAnimation("Hello " + name + " !", "luxisr.ttf", c.Writer)
}


