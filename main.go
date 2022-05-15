package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	//path URL "Anonymous function"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"messageA": "pong and pong",
		})
	})

	// todo Router GROUP VERSIONING
	v1 := router.Group("/v1")
	// Routing
	v1.GET("/hello", HelloHandler) //telah menggunakan groub

	router.GET("/books/:id/:tema", BooksHandler)
	router.GET("/query", QueryHandler)
	router.POST("/books", PostBookHandler)

	router.Run(":8888") //default port 8080
}

// Handler "function"
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":  "Welcome",
		"messageA": "Gin Gonic",
		"messageB": "Route Group path URL",
	})
}

// string ID `localhost:8888/books/20/anime`
func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	tema := c.Param("tema")

	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"tema": tema,
	})
}

// string Query `localhost:8888/query?title=hotgame&price=50000`
func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})

}

// Book Holder/contains information
type BookResponse struct {
	Title string      `json:"title" binding:"required"`        //validate
	Price json.Number `json:"price" binding:"required,number"` //json string/int yg penting angka
	//SubTitle string `json:"sub_title"`
}

func PostBookHandler(c *gin.Context) {
	// membuat object bookResponse
	var bookResponse BookResponse

	// if err := c.ShouldBindJSON(&bookResponse); err != nil { //cara satu
	err := c.ShouldBindJSON(&bookResponse)
	// Membuat Try and Catch
	// Jika error
	if err != nil {
		// Pesan Error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			//c.JSON(http.StatusBadRequest, errorMessage) //tanpa slice
			//return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

		//log.Fatal(err)
		// c.JSON(http.StatusBadRequest, err)
		// fmt.Println(err)
		// return
	}
	// jika success
	c.JSON(http.StatusOK, gin.H{
		"title": bookResponse.Title,
		"price": bookResponse.Price,
		//"sub_title": bookResponse.SubTitle,
	})
}
