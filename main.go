package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms
	router.MaxMultipartMemory = 8 << 30 //
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		// variable from the form
		priority := c.PostForm("priority")

		// listing uploading file
		for _, file := range files {
		log.Println(file.Filename) }

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, "/Users/xander/Desktop/"+filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()), priority)
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded! ", len(files)), priority)
	})
	router.Run(":8080")
}

// curl -X POST http://localhost:8080/upload -F "upload[]=@/Users/xander/Downloads/01.avi" -F "upload[]=@/Users/xander/Downloads/02.avi" -H "Content-Type: multipart/form-data"
// curl -X POST http://localhost:8080/upload -F "upload[]=@/Users/xander/Downloads/01.avi" -F "upload[]=@/Users/xander/Downloads/02.avi" -F "priority=8899" -H "Content-Type: multipart/form-data"
