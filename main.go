package main

import (
	"go-tutorial/rectangle"
	"log"

	"github.com/gin-gonic/gin"
)

type poligon interface {
	CalcArea()
}

func main() {
	/**
	var message []string
	var message map[string]string
	var message = "Hello"
	message := "Hello"
	//Map initialize
	product := make(map[string]string)
	//Literal concept
	product := map[string]string{
		"name": "zapato",
		"description": "zapato nuevo",
	}
	Print struct with keys
	fmt.Printf("%+v",a)
	*/
	rect := rectangle.NewRectangle(2, 3)
	rect.CalcArea()
	log.Println(rect.Area)
	draw(rect)

	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success my friend",
		})
	})
	r.Run()
}

func draw(figure poligon) {
	log.Printf("Dibuja")
}
