package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about the record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// initialize router and attached handlers to relative paths
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// attach router to a server on given address
	router.Run("localhost:8080")
}

// getAlbums endpoint responses with the albums slice serialized as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
