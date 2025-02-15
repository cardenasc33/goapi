package tools

import (
	log "github.com/sirupsen/logrus"
)


// Define types of data the database will return

// Database collections
type LoginDetails struct {
	AuthToken string
	Username string
}

type CoinDetails struct {
	Coins int64
	Username string
}

// Define methods for api
// Able to swap databases as long as 
// GetUserLoginDetails, GetUserCoins, and SetupDatabase are defined
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// function to return database
func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}