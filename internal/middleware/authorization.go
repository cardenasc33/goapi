package middleware

import (
	"github.com/cardenasc33/goapi/api"
	"github.com/cardenasc33/goapi/internal/tools"
	"errors"
	"net/http"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ResponseWriter used to construct a response to the caller
		// i.e. set the response body, header, status code
		// Request contains the information about the incoming request
		// i.e. Headers, payload, other information from http request that came in


		// Define logic for authorizing http request

		// Grab username param from request pointer 
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization") // auth token from header
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// instantiate a pointer to the database
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// Query database
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		// Return request error
		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// Calls next middleware in line or HandlerFunction
		// for the endpoint if there is no middleware left
		// e.g. Middlware -> next.ServeHTTP -> Middleware2 -> ... -> HandlerFunc
		// Actual: Authorization -> next.ServeHTTP -> GetCoinBalance
		
		next.ServeHTTP(w, r)
	})
}