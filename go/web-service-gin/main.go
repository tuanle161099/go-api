package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"` //to serialize and convert to json data when FE call
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

func albumById(c *gin.Context) {
	id := c.Param("id")
	album, error := getAlbumById(id)
	if error != nil {
		println(error)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found album"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func getAlbumById(id string) (*album, error) {
	for i, a := range albums {
		if a.ID == id {
			return &albums[i], nil
		}
	}
	return nil, errors.New("Not found album")
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", albumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
