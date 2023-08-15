package integration_test

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
	"userregistry/integration_test/helper"
	"userregistry/models"
	"userregistry/repo"
	mongo2 "userregistry/repo/mongo"
)

const (
	MongoDBUri = "mongodb://localhost:27017"
	Database   = "testdb"
	Collection = "users"
)

var pool *dockertest.Pool
var userTestRepo repo.UsersRepo
var dbClient *mongo.Client

func TestMain(m *testing.M) {
	var resource *dockertest.Resource
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic occurred", r)
		}

		if resource != nil {
			_ = pool.Purge(resource)
		}
	}()

	// Setup docker
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker %s", err)
	}

	pool = p

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// Setup docker options
	opts := dockertest.RunOptions{
		Repository:   "mongodb/mongodb-community-server",
		Tag:          "6.0-ubi8",
		ExposedPorts: []string{"27017"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"27017": {
				{HostIP: "0.0.0.0", HostPort: "27017"},
			},
		},
	}

	// Get resource
	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource %s", err)
	}

	// Start container and wait for it is ready
	if err := pool.Retry(func() error {
		var err error
		userTestRepo, err = mongo2.NewUsersRepo(MongoDBUri, Database, Collection)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		_ = pool.Purge(resource)
		log.Println("failed to init user repo", err)
	}

	// Populate data
	if err := helper.PopulateUsers(MongoDBUri, Database, Collection); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("Failed in user data population %s", err)
	}

	//run tests
	code := m.Run()

	//Clenup
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not cleanup docker container %s", err)
	}

	os.Exit(code)
}

func TestSaveUser(t *testing.T) {
	newUser := &models.User{
		Name:       "John",
		Password:   randomdata.Alphanumeric(10),
		ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
		Outputs: []string{
			randomdata.SillyName(),
			randomdata.SillyName(),
			randomdata.SillyName(),
		},
	}

	err := userTestRepo.Save(context.Background(), newUser)
	if err != nil {
		t.Errorf("User creation is failing")
	}
}

func TestGetUser(t *testing.T) {
	expectedUser := helper.Users[rand.Intn(len(helper.Users)-1)]
	user, err := userTestRepo.GetByUserName(context.Background(), expectedUser.Name)
	if err != nil {
		t.Errorf("error in finding user %v", err)
	}

	if user.Name != expectedUser.Name {
		t.Errorf("found username mismatch expected:%v, actual: %v, ", expectedUser.Name, user.Name)
	}

	if user.Password != expectedUser.Password {
		t.Errorf("found user password match expected:%v, actual: %v, ", expectedUser.Password, user.Password)
	}
}

func TestGetAll(t *testing.T) {
	users, err := userTestRepo.GetAll(context.Background())
	if err != nil {
		t.Errorf("error in finding user %v", err)
	}

	if len(helper.Users)+1 != len(users) {
		t.Errorf("length users not matching expected: %d, actual: %d", len(helper.Users), len(users))
	}
}
