package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"context"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mani関数が最初に実行される
func main() {

	StartWebServer()

}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc({ エンドポイント }, { レスポンス関数 }).Methods({ リクエストメソッド（複数可能） })
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/items", fetchAllItems).Methods("GET")
	// router.HandleFunc("/item/{id}", fetchSingleItem).Methods("GET")

	router.HandleFunc("/newitem", createItem).Methods("GET")
	// router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
	// router.HandleFunc("/item/{id}", updateItem).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

type dataType struct {
	Rid     string `bson:"rid"`
	Keyword string `bson:"keyword"`
}
func createItem(w http.ResponseWriter, r *http.Request) {	

	// 挿入するデータを作成
	data := dataType{
			Rid:     "俺のID",
			Keyword: "俺のキーワード",
	}

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	
	collection := client.Database("test").Collection("users")
	result, err := collection.InsertOne(context.Background(), data)
	print(err)
	print(result)
}

func fetchAllItems(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	
	collection := client.Database("test").Collection("users")
	cur, err := collection.Find(context.Background(), bson.M{
		"keyword": "俺のキーワード",
		})
	// 結果のカーソルをforで回して順番に結果を取得
	for cur.Next(context.Background()) {
		var ret dataType
		cur.Decode(&ret)
		fmt.Println(ret)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ret)
	}
}