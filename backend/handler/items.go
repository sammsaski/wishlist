/**
* This file is responsible for defining all API endpoints having the /items
* prefix, as we specified in the main handler file.
*/

package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sammsaski/wishlist-golang/backend/db"
	"github.com/sammsaski/wishlist-golang/backend/models"
)

// We will later use this variable for passing the itemID URL parameter
// across middlewares and request handlers using Go's context
var itemIDKey = "itemID"

// TODO: Add relevant documentation.
func items(router chi.Router) {
	router.Get("/", getAllItems)
	router.Post("/", createItem)
	router.Route("/{itemId}", func(router chi.Router) {
		router.Use(ItemContext)
		router.Get("/", getItem)
		router.Put("/", updateItem)
		router.Delete("/", deleteItem)
	})
}

/**
*	A chi middleware method to help us extract the `itemID` URL parameter
* 	from request URLs and use it in our code.
*
*	The middleware checks if `itemID` exists and is valid, and goes on to add
*	it to the request context.
*/
func ItemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "itemId")
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
			return
		}
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
		}
		ctx := context.WithValue(r.Context(), itemIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

/**
*	Decode a request into an instance of models.Item before sending it to
*	be saved in the database.
*/
func createItem(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddItem(item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

/**
*	Fetch all existing items in the database
*/
func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetAllItems()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, items); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

/**
*	View a specific item in the database
*/
func getItem(w http.ResponseWriter, r * http.Request) {
	itemID := r.Context().Value(itemIDKey).(int)
	item, err := dbInstance.GetItemById(itemID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

/**
*	Delete an item from the database
*/
func deleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.Context().Value(itemIDKey).(int)
	err := dbInstance.DeleteItem(itemID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

/** 
*	Update an item in the database
*/
func updateItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.Context().Value(itemIDKey).(int)
	itemData := models.Item{}
	if err := render.Bind(r, &itemData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	item, err := dbInstance.UpdateItem(itemID, itemData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))	
		}
		return
	}
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
