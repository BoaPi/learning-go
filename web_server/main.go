package main

import (
	"fmt"
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

	// GET Routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	// POST routes
	router.POST("/albums", postAlbums)

	// attach router to a server on given address
	router.Run("localhost:8080")
}

// getAlbums endpoint responses with the albums slice serialized as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// bind received JSON to newAlbum
	err := c.BindJSON(&newAlbum)

	if err != nil {
		fmt.Println("Bad Request, there is probably something wrong in the request")
		return
	}

	// add new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID returns the requested album from
// the slice of albums by matching the id  parameter and ID value of album
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over teh albums array to find the matching album
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// if album was not found in given slice of albums return "not found"
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
