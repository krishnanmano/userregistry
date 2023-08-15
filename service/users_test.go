package service

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
	"userregistry/models"
	"userregistry/repo/repofakes"
)

func TestCreate(t *testing.T) {
	fakeUsersRepo := repofakes.FakeUsersRepo{}
	fakeUsersRepo.SaveReturns(nil)

	userService := NewUserService(&fakeUsersRepo)
	err := userService.Create(context.Background(), &models.User{
		Name:       randomdata.FullName(randomdata.RandomGender),
		Password:   randomdata.Alphanumeric(10),
		ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
		Outputs: []string{
			randomdata.SillyName(),
			randomdata.SillyName(),
			randomdata.SillyName(),
		},
	})
	assert.Nil(t, err, "user creation error should be nil", err)
}

func TestGet(t *testing.T) {
	fakeUsersRepo := repofakes.FakeUsersRepo{}
	fakeUsersRepo.GetByUserNameReturns(&models.User{
		Name:       randomdata.FullName(randomdata.RandomGender),
		Password:   randomdata.Alphanumeric(10),
		ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
		Outputs: []string{
			randomdata.SillyName(),
			randomdata.SillyName(),
			randomdata.SillyName(),
		},
	}, nil)

	userService := NewUserService(&fakeUsersRepo)
	user, err := userService.Get(context.Background(), "test")
	assert.Nil(t, err, "get user error should be nil", err)
	assert.NotNil(t, user, "user should not be nil")
}

func TestUpdate(t *testing.T) {
	fakeUsersRepo := repofakes.FakeUsersRepo{}
	fakeUsersRepo.UpdateReturns(nil)

	userService := NewUserService(&fakeUsersRepo)
	username := randomdata.FullName(randomdata.RandomGender)
	err := userService.Update(context.Background(), &models.User{
		Name:       username,
		Password:   randomdata.Alphanumeric(10),
		ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
		Outputs: []string{
			randomdata.SillyName(),
			randomdata.SillyName(),
			randomdata.SillyName(),
		},
	}, username)
	assert.Nil(t, err, "update user error should be nil", err)
}

func TestGetAll(t *testing.T) {
	fakeUsersRepo := repofakes.FakeUsersRepo{}
	expectedUsers := []models.User{
		{
			Name:       randomdata.FullName(randomdata.RandomGender),
			Password:   randomdata.Alphanumeric(10),
			ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
			Outputs: []string{
				randomdata.SillyName(),
				randomdata.SillyName(),
				randomdata.SillyName(),
			},
		}, {
			Name:       randomdata.FullName(randomdata.RandomGender),
			Password:   randomdata.Alphanumeric(10),
			ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
			Outputs: []string{
				randomdata.SillyName(),
				randomdata.SillyName(),
				randomdata.SillyName(),
			},
		}, {
			Name:       randomdata.FullName(randomdata.RandomGender),
			Password:   randomdata.Alphanumeric(10),
			ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
			Outputs: []string{
				randomdata.SillyName(),
				randomdata.SillyName(),
				randomdata.SillyName(),
			},
		},
	}
	fakeUsersRepo.GetAllReturns(expectedUsers, nil)

	userService := NewUserService(&fakeUsersRepo)
	users, err := userService.GetAll(context.Background())
	assert.Nil(t, err, "get all users error should be nil", err)
	assert.Equal(t, len(expectedUsers), len(users), "get all users: length not matching")
}

func BenchmarkGetAll(b *testing.B) {
	fakeUsersRepo := repofakes.FakeUsersRepo{}
	expectedUsers := []models.User{
		{
			Name:       randomdata.FullName(randomdata.RandomGender),
			Password:   randomdata.Alphanumeric(10),
			ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
			Outputs: []string{
				randomdata.SillyName(),
				randomdata.SillyName(),
				randomdata.SillyName(),
			},
		}, {
			Name:       randomdata.FullName(randomdata.RandomGender),
			Password:   randomdata.Alphanumeric(10),
			ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
			Outputs: []string{
				randomdata.SillyName(),
				randomdata.SillyName(),
				randomdata.SillyName(),
			},
		}, {
			Name:       randomdata.FullName(randomdata.RandomGender),
			Password:   randomdata.Alphanumeric(10),
			ExpiryDate: time.Now().AddDate(rand.Intn(5), 0, 0).Unix(),
			Outputs: []string{
				randomdata.SillyName(),
				randomdata.SillyName(),
				randomdata.SillyName(),
			},
		},
	}
	fakeUsersRepo.GetAllReturns(expectedUsers, nil)

	userService := NewUserService(&fakeUsersRepo)
	users, err := userService.GetAll(context.Background())
	assert.Nil(b, err, "get all users error should be nil", err)
	assert.Equal(b, len(expectedUsers), len(users), "get all users: length not matching")
}
