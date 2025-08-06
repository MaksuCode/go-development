package main

import (
	"math"
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
	var foundAlbums []album
	minPrice := 0.0
	maxPrice := math.MaxFloat64

	if minStr := c.Query("min"); minStr != "" {
		if val, err := strconv.ParseFloat(minStr, 64); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid min price",
				"error":   err.Error(),
			})
			return
		} else {
			minPrice = val
		}
	}

	if maxStr := c.Query("max"); maxStr != "" {
		if val, err := strconv.ParseFloat(maxStr, 64); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid max price",
				"error":   err.Error(),
			})
			return
		} else {
			maxPrice = val
		}
	}

	if minPrice > maxPrice {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "min price can not be bigger than max price",
		})
		return
	}

	for _, album := range albums {
		if album.Price <= maxPrice && album.Price >= minPrice {
			foundAlbums = append(foundAlbums, album)
		}
	}

	if len(foundAlbums) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no album found with the given min-max range"})
		return
	}
	c.IndentedJSON(http.StatusOK, foundAlbums)
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

func getAlbumByTitle(c *gin.Context) {
	title := c.Param("title")

	for _, album := range albums {
		if album.Title == title {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/id/:id", getAlbumByID)
	router.GET("/albums/title/:title", getAlbumByTitle)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")
}
