package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// mani関数が最初に実行される
func main() {
	StartWebServer()
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		panic(err.Error())
	}
	str := "Sample Message"
	tmpl.Execute(w, str)
}

func StartWebServer() {
	fmt.Println("Rest API with Mux Routers")
	dir, _ := os.Getwd()
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/items", fetchAllItems)
	http.HandleFunc("/newitem", createItem)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
	return
}

type dataType struct {
	id   string   `json:"id"`
	data []string `json:"data"`
}

func createItem(w http.ResponseWriter, r *http.Request) {

	// 挿入するデータを作成
	// data := dataType{
	// 	id:      "2",
	// 	data: [],
	// }

	// clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	panic(err)
	// }

	// collection := client.Database("test").Collection("users")
	// result, err := collection.InsertOne(context.Background(), data)
	// print(err)
	// print(result)
}

type FindOneRequest struct {
	id string `json:"id"`
}

func fetchAllItems(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	var doc bson.Raw
	filter := bson.D{{"id", "1"}}
	collection := client.Database("test").Collection("users")
	err = collection.FindOne(context.Background(), filter).Decode(&doc)
	fmt.Println(doc.String())
	var docBsonD bson.D
	err = bson.Unmarshal(doc, &docBsonD)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(docBsonD)
}
