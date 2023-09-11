package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("POST request received"))
		var result Item
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		result.Name = r.FormValue("name")
		_, err := collection.InsertOne(context.Background(), result)
		if err != nil {
			http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
			return
		}
	
		w.WriteHeader(http.StatusCreated)

	}else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusFound)
}

func handleRequests() {
	http.HandleFunc("/api", homePage)
	log.Fatal(http.ListenAndServe(":4002", nil))
}

var client *mongo.Client
var collection *mongo.Collection

func init() {
	// Connect to MongoDB.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Update with your MongoDB URI.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection to MongoDB.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set up a collection to store items.
	collection = client.Database("your_db_name").Collection("items") // Update with your DB name and collection name.
}

func main() {
	handleRequests()
	fmt.Println("start")
	handleRequests()
}
