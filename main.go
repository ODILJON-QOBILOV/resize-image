package main

import (
	"bytes"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/resize", ResizeImageAPI)

	r.Run()
}

func ResizeImageAPI(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	
	width, _ := strconv.Atoi(c.DefaultPostForm("width", "200"))
	height, _ := strconv.Atoi(c.DefaultPostForm("height", "200"))

	img, err := imaging.Decode(file)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resizedImage := imaging.Resize(img, width, height, imaging.Lanczos)

	var buf bytes.Buffer
	err = imaging.Encode(&buf, resizedImage, imaging.PNG)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Data(200, "image/png", buf.Bytes())
}