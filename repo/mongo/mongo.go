package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"userregistry/common/errors"
	"userregistry/models"
	"userregistry/repo"
	models2 "userregistry/repo/models"
)

type usersRepo struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewUsersRepo(uri, database, collection string) (repo.UsersRepo, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	coll := client.Database(database).Collection(collection)
	return &usersRepo{
		mongoClient: client,
		collection:  coll,
	}, nil
}

func (u *usersRepo) Save(ctx context.Context, user *models.User) error {
	_, err := u.collection.InsertOne(ctx, user)
	return err
}

func (u *usersRepo) Update(ctx context.Context, user *models.User, userName string) error {
	oldUser := models2.User{}
	filter := bson.D{{"username", userName}}
	err := u.collection.FindOne(ctx, filter).Decode(&oldUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.ErrUserNotFound
		}
		return err
	}

	filter = bson.D{{"_id", oldUser.ID}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", user}}
	_, err = u.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func (u *usersRepo) GetByUserName(ctx context.Context, userName string) (*models.User, error) {
	user := models.User{}
	filter := bson.D{{"username", userName}}
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *usersRepo) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	filter := bson.D{}
	cur, err := u.collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)

	return users, nil
}

func (u *usersRepo) Close() error {
	if err := u.mongoClient.Disconnect(context.Background()); err != nil {
		return err
	}
	return nil
}
