package handlers

import (
	"github.com/cardenasc33/goapi/api"
	"github.com/cardenasc33/goapi/internal/tools"
	"encoding/json"
	"net/http"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	// Grab username from the parameters passed in
	
	// Decode parameters
	var params = api.CoinBalanceParams{} 
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// Grab params from URL and set them to the values in the struct
	// e.g. Grab username in URL and put it in the username
	// field in the struct
	err = decoder.Decode(&params, r.URL.Query())  

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// instantiate a db interface
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase() 
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Call GetUserCoins method
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Set value to the response struct
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}

	// Write the response struct to the response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}