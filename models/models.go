package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"golang-review-phone/database"
)

type Review struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Comments    []Comment          `json:"comments,omitempty" bson:"comments,omitempty"`
}

type Comment struct {
	UserID  string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Role     string             `json:"role,omitempty" bson:"role,omitempty"`
}

type LoginCredentials struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

var collection *mongo.Collection

func InitModels() {
	collection = database.GetCollection("reviews")
}

func CreateReview(review Review) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, review)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func GetReviews() ([]Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []Review
	err = cursor.All(ctx, &reviews)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return reviews, nil
}

func GetReviewByID(id string) (*Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var review Review
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&review)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &review, nil
}

func PostComment(id string, comment Comment) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}

	update := bson.M{
		"$push": bson.M{"comments": comment},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}

	return result.UpsertedID.(primitive.ObjectID), nil
}

func RegisterUser(user User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func GetUserByUsername(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
