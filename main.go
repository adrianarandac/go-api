package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct{
	ID 			string 		`json:"id"`
	Title 		string 		`json:"title"`
	Author 		string 		`json:"author"`
	Quantity	 int 		`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Don Quijote", Author: "Cervantes", Quantity: 4},
	{ID: "2", Title: "1984", Author: "Orson Welles", Quantity: 1},
	{ID: "3", Title: "Holy Bible", Author: "Lotta People", Quantity: 69},
}

func getBooks (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById (c *gin.Context)  {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book) 
}

//CHECKOUT
func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H {"message": "book out of stock"})
		return
	}
	book.Quantity--
	c.IndentedJSON(http.StatusOK, book) 
}

//CHECKOUT
func returnBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity++
	c.IndentedJSON(http.StatusOK, book) 
}



func getBookById (id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func createBook (c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return	
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main () {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}