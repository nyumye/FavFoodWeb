package main

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collections map[string]*mongo.Collection
var databaseName = "favfoodweb"
var collectionNames = []string{"foods"}

type foodDataModel struct {
	ID        string `json:"_id" bson:"_id, omitempty"`
	Name      string `json:"name" bson:"name"`
	BriefDesc string `json:"brief_desc" bson:"brief_desc"`
	MainDesc  string `json:"main_desc" bson:"main_desc"`
	ImageUri  string `json:"image_uri" bson:"image_uri"`
}

func prepareMongoClient() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

}

func prepareMongoCollection() error {
	if client == nil {
		return errors.New(`Call "prepareMongoClient" method before "prepareMongoCollection" method`)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	collections = make(map[string]*mongo.Collection)
	for _, colName := range collectionNames {
		collections[colName] = client.Database(databaseName).Collection(colName)
	}
	return nil
}

// register initial value to collections[any]
func registerInitialFoodData() {
	initialFoodDatas := makeInitialFoodData()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// opts := options.Update().SetUpsert(true)

	for _, initialFoodData := range initialFoodDatas {
		//collections["foods"].UpdateOne(ctx, bson.D{{Key: "_id", Value: initialFoodData.ID}}, initialFoodData, opts)
		collections["foods"].InsertOne(ctx, initialFoodData, options.InsertOne())
	}
}

// make initial food datas then return []foodDatamodel
func makeInitialFoodData() []foodDataModel {
	return []foodDataModel{
		foodDataModel{
			ID:        "0",
			Name:      "ドーナツ",
			BriefDesc: "おいしいドーナツ",
			MainDesc:  "おいしいけどカロリーおよび糖質が気になる。",
			ImageUri:  "./image/donut.png",
		},
		foodDataModel{
			ID:        "1",
			Name:      "焼きもち",
			BriefDesc: "あの日夢見たビジュアル系焼きもち",
			MainDesc:  "焼いた餅。こんな感じに膨らむことは無いし、膨らんだ面に焼き色が付くことは無いがこうしたほうがおいしそうなのでそうした。ぜんざいに入れるとおいしいのだが、この前1kgのあんこ買ってきたら毎日のように食べてしまったし、もちも食べる度に二個食べてしまったので危ないなと思ってそれ以来買っていない。",
			ImageUri:  "./image/rice-cake.png",
		},
	}
}

func findAllFoodDocument() []foodDataModel {
	var ctx context.Context
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collections["foods"].Find(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	cancel()

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var results []foodDataModel
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	cancel()
	return results
}

func prepareDatabase() {
	// favfoodweb -> foods collection
	prepareMongoClient()
	if err := prepareMongoCollection(); err != nil {
		log.Fatal(err)
	}

	registerInitialFoodData()

	// foods := findAllFoodDocument()
	// for _, food := range foods {
	// 	fmt.Println(food.Name)
	// }

}
