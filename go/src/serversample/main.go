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

	// router.HandleFunc("/item", createItem).Methods("POST")
	// router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
	// router.HandleFunc("/item/{id}", updateItem).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
type ItemParams struct {
	Id           string    `json:"id"`
	JanCode      string    `json:"jan_code,omitempty"`
	ItemName     string    `json:"item_name,omitempty"`
	Price        int       `json:"price,omitempty"`
	CategoryId   int       `json:"category_id,omitempty"`
	SeriesId     int       `json:"series_id,omitempty"`
	Stock        int       `json:"stock,omitempty"`
	Discontinued bool      `json:"discontinued"`
	ReleaseDate  time.Time `json:"release_date,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

var items []*ItemParams

func fetchAllItems(w http.ResponseWriter, r *http.Request) {
	type dataType struct {
		Rid     string `bson:"rid"`
		Keyword string `bson:"keyword"`
	}

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

	
	cur, err := collection.Find(context.Background(), bson.M{
		"keyword": "俺のキーワード",
		})
	// 結果のカーソルをforで回して順番に結果を取得
	for cur.Next(context.Background()) {
		var ret dataType
		cur.Decode(&ret)
		fmt.Println(ret)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func init() {
	items = []*ItemParams{
			&ItemParams{
					Id:           "1",
					JanCode:      "327390283080",
					ItemName:     "item_1",
					Price:        2500,
					CategoryId:   1,
					SeriesId:     1,
					Stock:        100,
					Discontinued: false,
					ReleaseDate:  time.Now(),
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    time.Now(),
			},
			&ItemParams{
					Id:           "2",
					JanCode:      "3273902878656",
					ItemName:     "item_2",
					Price:        1200,
					CategoryId:   2,
					SeriesId:     2,
					Stock:        200,
					Discontinued: false,
					ReleaseDate:  time.Now(),
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    time.Now(),
			},
	}
}