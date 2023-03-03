package repository

import (
	"GoGraphQlJwt/graph/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoRepository interface {
	Save(video *model.Video)
	FindAll() []*model.Video
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE   = "gographql"
	COLLECTION = "videos"
)

func New() VideoRepository {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions = clientOptions.SetMaxPoolSize(50)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database")
	return &database{
		client: dbClient,
	}
}

func (db *database) Save(video *model.Video) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) FindAll() []*model.Video {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Video
	for cursor.Next(context.TODO()) {
		var v *model.Video
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}
