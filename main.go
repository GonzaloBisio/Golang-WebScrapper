package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"name"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "21", Artist: "Adele", Year: 2011},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2022},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
func main() {
	fmt.Println("Starting server...")
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumById)
	router.Run("localhost:8080")
}
