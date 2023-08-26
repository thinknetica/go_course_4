package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "data"      // имя учебной БД
	collectionName = "languages" // имя коллекции в учебной БД
)

type lang struct {
	ID   int
	Name string
}

func main() {
	// подключение к СУБД MongoDB
	mongoOpts := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}

	// не забываем закрывать ресурсы
	defer client.Disconnect(context.Background())

	// проверка связи с БД
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// вставка пары документов в БД
	data := []lang{
		{ID: 1, Name: "Go"},
		{ID: 2, Name: "JavaScript"},
	}
	err = insertDocs(client, data)
	if err != nil {
		log.Fatal(err)
	}

	// получение документов из БД
	docs, err := docs(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(docs)
}

// insertDocs вставляет в БД моссив документов.
func insertDocs(c *mongo.Client, data []lang) error {
	collection := c.Database(databaseName).Collection(collectionName)
	for _, doc := range data {
		_, err := collection.InsertOne(context.Background(), doc)
		if err != nil {
			return err
		}
	}
	return nil
}

// docs возвращает все документы из БД.
func docs(c *mongo.Client) ([]lang, error) {
	collection := c.Database(databaseName).Collection(collectionName)
	filter := bson.D{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var data []lang
	for cur.Next(context.Background()) {
		var l lang
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}

	return data, cur.Err()
}
