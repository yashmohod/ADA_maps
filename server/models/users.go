package models

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id           uuid.UUID `bson:"_id"`
	Email        string    `bson:"email,omitempty"`
	PasswordHash string    `bson:"passwordhash,omitempty"`
	FirstName    string    `bson:"firstname,omitempty"`
	LastName     string    `bson:"lastname,omitempty"`
	SessionToken string    `bson:"sessionToken,omitempty"`
	CsrfToken    string    `bson:"csrfToken,omitempty"`
}

func AddTokens(userId uuid.UUID, sessionToken string, csrfToken string) error {
	collection := client.Database(db).Collection("user")

	// id, _ := bson.ObjectIDFromHex(userId)
	filter := bson.M{"_id": userId}

	update := bson.M{"$set": bson.M{"sessionToken": sessionToken, "csrfToken": csrfToken}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func RegisterUser(user User) error {
	collection := client.Database(db).Collection("user")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetUser(email string) (User, error) {
	collection := client.Database(db).Collection("user")
	log.Println(email)
	filter := bson.M{"email": email}
	var result User
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Println(err)
		return User{}, err
	}
	return result, nil
}

func DeleteUser(email string) error {
	collection := client.Database(db).Collection("user")
	filter := bson.M{"email": email}
	// Deletes the first document that has a "title" value of "Twilight"
	_, err := collection.DeleteOne(context.TODO(), filter)
	// Prints a message if any errors occur during the operation
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserInfo(replacement User) error {
	collection := client.Database(db).Collection("user")

	filter := bson.M{"_id": replacement.Id}

	// Replaces the first document that matches the filter with a new document
	_, err := collection.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
