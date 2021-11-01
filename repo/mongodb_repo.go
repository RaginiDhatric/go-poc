package repo

import (
	"bookusecase/entity"
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
	CONNECTIONSTRING = "mongodb://mongodb:27017"
	DB               = "books_db"
	BOOKS            = "books_collection"
)

type mongoDBRepo struct{}

func NewMongoDBRepo() BookRepository {
	return &mongoDBRepo{}
}

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func (*mongoDBRepo) Get() ([]*entity.Book, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}}
	var books []*entity.Book
	
	client, err := GetMongoClient()
	if err != nil {
		return books, err
	}
	
	collection := client.Database(DB).Collection(BOOKS)
	
	cursor, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return books, findError
	}
	//Map result to slice
	for cursor.Next(context.TODO()) {
		var t *entity.Book
		err := cursor.Decode(&t)
		if err != nil {
			return books, err
		}
		books = append(books, t)
	}

	cursor.Close(context.TODO())
	if len(books) == 0 {
		return books, mongo.ErrNoDocuments
	}
	return books, nil
}

func (*mongoDBRepo) Create(book *entity.Book) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(DB).Collection(BOOKS)
	_, err = collection.InsertOne(context.TODO(), book)
	if err != nil {
		return err
	}
	return nil
}