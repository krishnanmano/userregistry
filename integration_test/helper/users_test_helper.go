package helper

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"math/rand"
	"time"
	"userregistry/models"
)

var Users []models.User

func PopulateUsers(uri, database, collection string) error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	col := client.Database(database).Collection(collection)

	for i := 1; i <= 10; i++ {
		user := models.User{
			Name:       randomdata.FullName(randomdata.RandomGender),
			Password:   randomdata.Alphanumeric(10),
			ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
			Outputs: []string{
				randomdata.SillyName(),
				randomdata.SillyName(),
				randomdata.SillyName(),
			},
		}

		_, err := col.InsertOne(context.Background(), user)
		if err != nil {
			return err
		}

		Users = append(Users, user)
	}
	return nil
}
