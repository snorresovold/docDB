package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/snorresovold/docDB/filesys"
)

func createCollection(c *gin.Context) {
	name := c.Param("name")
	filesys.CreateCollection(name)
	c.IndentedJSON(http.StatusOK, name)
}

func createDocument(c *gin.Context) {
	collection := c.Param("collection")
	filesys.CreateDocument(collection)
	c.IndentedJSON(http.StatusOK, collection)
}

func writeToDocument(c *gin.Context) {
	collection := c.Param("collection")
	document := c.Param("document")
	// Parse JSON data from the request body into a map
	var data map[string]interface{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err, "error") //print the error on the console
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err, "error") //print the error on the console
		return
	}
	// Call the WriteToDocument function with the document name and data
	filesys.WriteToDocument(collection, document, data)
	c.IndentedJSON(http.StatusOK, document)
}

func deleteCollection(c *gin.Context) {
	collection := c.Param("collection")
	filesys.DeleteCollection(collection)
	c.IndentedJSON(http.StatusOK, collection)
}

func deleteDocument(c *gin.Context) {
	collection := c.Param("collection")
	document := c.Param("document")
	filesys.DeleteDocument(collection + "/" + document)
	c.IndentedJSON(http.StatusOK, collection+document)
}

func getCollection(c *gin.Context) {
	collection := c.Param("collection")
	out := filesys.GetCollection(collection)
	c.IndentedJSON(http.StatusOK, out)
}
func getDocument(c *gin.Context) {
	collection := c.Param("collection")
	document := c.Param("document")
	out := filesys.GetDocument(collection, document)
	c.IndentedJSON(http.StatusOK, out)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/getCollection/:collection", getCollection)
	r.GET("/getDocument/:collection/:document", getDocument)
	r.POST("/createCollection/:name", createCollection)
	r.POST("/createDocument/:collection", createDocument)
	r.POST("/writeToDocument/:collection/:document", writeToDocument)
	r.DELETE("/deleteCollection/:collection", deleteCollection)
	r.DELETE("/deleteDocument/:collection/:document", deleteDocument)
	r.Run("localhost:8080")
}