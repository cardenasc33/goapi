package api

import (
	"encoding/json"
	"net/http"
)

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

	// Name of target
	Name string

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

// Writes an error response to the http response writer
// returns an error response to the person who called the endpoint
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	// Write to response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// Wrapper to provide different types of error responses
var (

	// Specific error response to tell user to fix their request
	// e.g. incorrect or missing username or password
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	// Interal error, e.g. bug in the code, respond with generic error message
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)