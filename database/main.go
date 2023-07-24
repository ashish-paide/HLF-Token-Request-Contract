package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"encoding/json"
)

// Define a struct to represent the JSON payload


type Request struct{
	Id int `json:"id"`
	Usernname string `json:"username"`
	Amount string `json:"amount"`
	Remark string `json:"remark"`
}



var db = Create_Database("db")

func main() {


    defer db.db.Close()

	// Create a new Gin router
	router := gin.Default()

	// POST /create/{key}
	router.POST("/create/:key", createHandler)
	router.GET("/get/:key", getHandler)
	router.GET("/getall", getAllHandler)
	router.GET("/delete/:key" , deleteHandler)

	// Start the server on port 8080
	log.Println("Server listening on port 18080...")
	router.Run(":18080")
}

// Handler for the POST endpoint /create/{key}
func createHandler(c *gin.Context) {
	// Extract the key from the URL path
	key := (c.Param("key"))

	// Decode the JSON payload from the request body
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Insert the data into Jhuno DB with the key
	db.Set(key, request)


	// Send a success response
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Data inserted successfully"})
}

// Handler for the GET endpoint /get/{key}
func getHandler(c *gin.Context) {
	// Extract the key from the URL path
	key := c.Param("key")

	// Read the data from Jhuno DB based on the key
	data , err:= db.Get(key)
	if err != nil {
		fmt.Println("unable to get from key:", key)
	}

	// Send the data as a JSON response
	c.JSON(http.StatusOK , gin.H{ "message": data})
}

func getAllHandler(c *gin.Context) {
	data := db.Getall()
	
	var jsonData interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing JSON data")
		return
	}
	fmt.Println(jsonData)
	c.JSON(http.StatusOK, gin.H{"status": "fetched successfully", "message": jsonData})
}


func deleteHandler(c *gin.Context){
	key := c.Param("key")
	
	data , err:= db.Get(key)
	if err != nil {
		fmt.Println("unable to get from key:", key)
	}

	err = db.Delete(key)
	if err != nil {
		fmt.Println("cannot delete key: ", err)
		c.JSON(http.StatusOK, gin.H{"status": "cannot delete key", "message": err})
	}

	

	c.JSON(http.StatusOK, gin.H{"status": "deleted successfully", "message": data })

}

