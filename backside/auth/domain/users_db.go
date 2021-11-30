package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sk8.town/backside/auth/config"
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/users_db.go -package=mocks sk8.town/backside/auth/domain UsersRepository
type UsersRepository interface {
	Add(*User) *errs.AppError
	Get(email string) (*User, *errs.AppError)
}

const (
	usersCollection = "users"
)

type UsersDb struct {
}

func (db UsersDb) Add(user *User) *errs.AppError {
	client, err := GetMongoClient()
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	collection := client.Database(config.GetConfig().DbName).Collection(usersCollection)

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	return nil
}

func (db UsersDb) Get(email string) (*User, *errs.AppError) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	collection := client.Database(config.GetConfig().DbName).Collection(usersCollection)

	result := User{}
	filter := bson.D{primitive.E{Key: "email", Value: email}}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, nil
	}

	return &result, nil
}

func NewUsersDb() UsersDb {
	return UsersDb{}
}
