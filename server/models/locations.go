package models

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Location struct {
	Id        uuid.UUID `bson:"_id"`
	Latitude  float64   `bson:"latitude,omitempty"`
	Longitude float64   `bson:"longitude,omitempty"`
	Name      string    `bson:"name,omitempty"`
	EntryBy   uuid.UUID `bson:"entryby"`
}

func AddLocation(location Location) error {
	collection := client.Database(db).Collection("location")
	_, err := collection.InsertOne(context.TODO(), location)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetLocations() ([]Location, error) {
	log.Println("got here!")
	collection := client.Database(db).Collection("location")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}

	var results []Location
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
		return []Location{}, err
	}
	return results, nil
}

func GetLocation(locationId string) (Location, error) {
	collection := client.Database(db).Collection("location")
	uid, er := uuid.Parse(locationId)
	if er != nil {
		log.Println(er)
		return Location{}, er
	}
	filter := bson.M{"_id": uid}
	var result Location
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Println(err)
		return Location{}, err
	}
	return result, nil
}

func DelteleLocation(locationId uuid.UUID) error {
	collection := client.Database(db).Collection("location")
	filter := bson.M{"_id": locationId}

	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateLocation(replacement Location) error {
	collection := client.Database(db).Collection("location")

	filter := bson.M{"_id": replacement.Id}

	// Replaces the first document that matches the filter with a new document
	_, err := collection.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
