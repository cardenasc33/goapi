package api

// import (
// 	"encoding/json"
// 	"net/http"
// )

// PARAMETERS & RESPONSES OF ENDPOINT

// Coin Balance Params Struct:
// Represents parameters that the API endpoint will take
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response:
// Outlines successful response from the server containing
// the status code and account balance
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response Struct:
// Represents response returned when an error occurs
type Error struct {
	//Error code
	Code int

	// Error message
	Message string
}
