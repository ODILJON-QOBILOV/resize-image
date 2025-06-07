package controllers

import (
	"bytes"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

// ResizeImageAPI godoc
// @Summary Resize an uploaded image
// @Description Accepts an image file and optional width and height, returns resized image in PNG format
// @Tags image
// @Accept multipart/form-data
// @Produce image/png
// @Param image formData file true "Image file to resize"
// @Param width formData int false "Width" default(200)
// @Param height formData int false "Height" default(200)
// @Success 200 {file} binary "Resized image PNG"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /resize [post]
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
