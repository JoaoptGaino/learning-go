package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", createAlbums)
	router.Run("localhost:8080")
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Kill em all", Artist: "Metallica", Price: 56.99},
	{ID: "2", Title: "Ride the lightning", Artist: "Metallica", Price: 59.99},
	{ID: "3", Title: "The number of the beast", Artist: "Iron Maiden", Price: 59.99},
	{ID: "4", Title: "Powerslave", Artist: "Iron Maiden", Price: 59.99},
	{ID: "5", Title: "Peace sells", Artist: "Megadeth", Price: 59.99},
	{ID: "6", Title: "Rust in peace", Artist: "Megadeth", Price: 59.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func createAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
