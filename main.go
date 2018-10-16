package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lwnmengjing/queueMail/form"
	"github.com/lwnmengjing/queueMail/mail"
	"net/http"
)

var (
	listen         = flag.String("listen", ":8080", "Server listen port(example :12312)")
	filepath       = flag.String("filepath", "./attachments", "Attachment's file path")
	MessageChannel = make(chan form.MessageMail)
)

func init() {
	go Send(MessageChannel)
}

func main() {
	r := gin.Default()
	r.POST("/message", messageHandle)
	r.POST("/upload", func(c *gin.Context) {
		// Multipart form
		f, _ := c.MultipartForm()
		files := f.File["upload[]"]
		var filepathes []string

		for _, file := range files {
			// Upload the file to specific dst.
			err := c.SaveUploadedFile(file, *filepath+"/"+file.Filename)
			filepathes = append(filepathes, *filepath+"/"+file.Filename)
			if err != nil {
				c.JSON(http.StatusPartialContent, gin.H{
					"message": fmt.Sprintf("File(%s) save failed", file.Filename),
					"files":   filepathes,
				})
				c.Abort()
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "All files is Uploaded!",
			"files":   filepathes,
		})
		c.Abort()
	})
	r.Run(*listen)
}

func messageHandle(c *gin.Context) {
	var message form.MessageMail
	if c.BindJSON(&message) != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid form", "form": message})
		c.Abort()
		return
	}
	go func() {
		MessageChannel <- message
	}()
	c.JSON(http.StatusAccepted, gin.H{
		"meesage": "The message has been queued and sent later!",
	})
	c.Abort()
}

func Send(input chan form.MessageMail) {
	for {
		select {
		case message := <-input:
			mail.SendMail(message)
		}
	}
}
