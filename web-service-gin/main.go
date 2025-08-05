package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func getAlbumWithinPriceRange(c * gin.Context) []album {
	min, err := strconv.ParseFloat(c.Query("min"), 64)
	if err != nil {
		min = 0
	}
	max, err := strconv.ParseFloat(c.Query("max"), 64)
	if err != nil {
		max = 10000
	}

	var foundAlbums []album

	for _, album := range albums {
		if album.Price < max && album.Price > min {
			foundAlbums = append(foundAlbums, album)
		}
	}
	return foundAlbums;
}

func getAlbumByTitle(c *gin.Context) {
	title := c.Param("title")

	for _, album := range albums {
		if album.Title == title{
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/id/:id", getAlbumByID)
	router.GET("/albums/title/:title", getAlbumByTitle)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")
}
