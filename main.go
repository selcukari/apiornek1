package main

import (
	"fmt"
	"net/http"
	"web-api/handlers"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	age  int
}

func (p *Person) Greet(name string) {
	p.Name = name
	fmt.Printf("Hello, %s! My name is %s and I am %d years old.\n", name, p.Name, p.age)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.GET("/contents", func(c *gin.Context) {
		handlers.GetAllContents(c.Writer, c.Request)
	})

	router.GET("/contents/:id", func(c *gin.Context) {
		handlers.GetContentID(c.Writer, c.Request)
	})

	router.Run("localhost:8080")
}
