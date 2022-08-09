/**
* The code here is responsible for interacting directly with our database.
* This way, the database engine is properly separated from the rest of the
* application.
*/

package db

import (
	"database/sql"
	"github.com/sammsaski/wishlist-golang/backend/models"
)

/**
*	Return all items in the database.
*/
func (db Database) GetAllItems() (*models.ItemList, error) {
	list := &models.ItemList{}
	items, err := db.Conn.Query("SELECT * FROM wishlist ORDER BY ID DESC;")
	if err != nil {
		return list, err
	}
	for items.Next() {
		var item models.Item
		err := items.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

/**
*	Add an item to the database.
*/
func (db Database) AddItem(item *models.Item) error {
	var id int
	query := `INSERT INTO wishlist (name, price) VALUES ($1, $2) RETURNING id;`
	err := db.Conn.QueryRow(query, item.Name, item.Price).Scan(&id)
	if err != nil {
		return err
	}
	item.ID = id
	return nil
}

/**
*	Retrieve a specific item (by UID) from the database.
*/
func (db Database) GetItemById(itemId int) (models.Item, error) {
	item := models.Item{}
	query := `SELECT * FROM wishlist WHERE id = $1;`
	queryResult := db.Conn.QueryRow(query, itemId)
	switch err := queryResult.Scan(&item.ID, &item.Name, &item.Price); err {
	case sql.ErrNoRows:
		return item, ErrNoMatch
	default:
		return item, err
	}
}

/**
*	Remove a specific item (by UID) from the database.
*/
func (db Database) DeleteItem(itemId int) error {
	query := `DELETE FROM wishlist WHERE id = $1;`
	_, err := db.Conn.Exec(query, itemId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

/**
*	Update a specific item (by UID) from the database.
*/
func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
	item := models.Item{}
	query := `UPDATE wishlist SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price;`
	err := db.Conn.QueryRow(query, itemData.Name, itemData.Price, itemId).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, ErrNoMatch
		}
		return item, err
	}
	return item, nil
}
