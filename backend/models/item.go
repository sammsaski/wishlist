/**
* Contains Go structs that can be bounded to database objects or
* transformed into their equivalent JSON format.
*/

package models

import (
	"fmt"
	"net/http"
)

/**
*	An item in our wishlist
*/
type Item struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

type ItemList struct {
	Items []Item `json:"items"`
}

func (i *Item) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
