package main

import (
	"fmt"
	"net/http"
	"strings"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mani関数が最初に実行される
func main() {
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
}