package database

import (
	"context"
	"log"
	"time"

	"gihub.com/aalekh/gqlmdb/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dialter = "mongodb://127.0.0.1:27017/gsqldb"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {

	client, err := mongo.NewClient(options.Client().ApplyURI(dialter))

	if err != nil {
		log.Println("Failed To create Database", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}

}

func (db *DB) CreateJobListing(input model.CreateJobListingInput) *model.JobListing {
	jobcollection := db.client.Database("gsqldb").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserted, err := jobcollection.InsertOne(ctx, model.CreateJobListingInput{
		Title:       input.Title,
		Description: input.Description,
		Company:     input.Company,
		URL:         input.URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	inseertedid := inserted.InsertedID.(primitive.ObjectID).Hex()
	createdresponse := model.JobListing{
		ID:          inseertedid,
		Title:       input.Title,
		Description: input.Description,
		Company:     input.Company,
		URL:         input.URL,
	}
	return &createdresponse
}

func (db *DB) UpdateJobListing(id string, input model.UpdateJobListingInput) *model.JobListing {
	jobcollection := db.client.Database("gsqldb").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updatejob := bson.M{}

	if input.Title != nil {
		updatejob["title"] = input.Title
	}

	if input.Description != nil {
		updatejob["description"] = input.Description
	}
	if input.URL != nil {
		updatejob["url"] = input.URL
	}

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updatejob}

	results := jobcollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var joblisting model.JobListing

	err := results.Decode(&joblisting)
	if err != nil {
		log.Fatal(err)
	}

	return &joblisting
}

func (db *DB) DeleteJobListing(id string) *model.DeleteJobResponse {
	jobcollection := db.client.Database("gsqldb").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	_, err := jobcollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return &model.DeleteJobResponse{DeletedJobID: id}
}

func (db *DB) Getjobs() []*model.JobListing {
	jobcollection := db.client.Database("gsqldb").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var joblisting []*model.JobListing

	cursor, err := jobcollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &joblisting); err != nil {
		panic(err)
	}

	return joblisting
}

func (db *DB) Getjob(id string) *model.JobListing {

	jobcollection := db.client.Database("gsqldb").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	fileter := bson.M{"_id": _id}
	var joblisting model.JobListing
	err := jobcollection.FindOne(ctx, fileter).Decode(&joblisting)
	if err != nil {
		log.Fatal(err)
	}

	return &joblisting
}
