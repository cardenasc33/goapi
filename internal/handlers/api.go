package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/cardenasc33/goapi/internal/middleware"
)

// Capital H to tell compiler this function can be imported
// in other packages
func Handler(r *chi.Mux) {
	// Add middleware to a route using r.Use()
	// Stripe slashes ignores trailing slashes at the end
	r.Use(chimiddle.StripSlashes) // global middleware, applied to all endpoints

	// Route Setup:
	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)  // check if user is authrized to access data

		router.Get("/coins", GetCoinBalance)
	})

}