/**
* Here, we provide the functionality for an API handler that will route
* HTTP requests that hit the /items endpoint to their corresponding 
* response (GET, PUT, DELETE).
*/

package handler

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sammsaski/wishlist-golang/backend/db"
)

var dbInstance db.Database

/**
*	Create an instance of the handler.
*/
func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/items", items)
	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}

