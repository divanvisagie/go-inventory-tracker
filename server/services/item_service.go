package services

import (
	"fmt"

	"github.com/divanvisagie/go-inventory-tracker/server/models"
	"github.com/divanvisagie/go-inventory-tracker/server/restapi/operations/items"
	"github.com/go-openapi/swag"
	"github.com/go-pg/pg"
)

func logDBIssue(err error) {
	if err != nil {
		fmt.Errorf(">> Database Error %s", err.Error)
	}
}

func getDbConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "secret",
		Addr:     "postgresdb:5432",
	})
	return db
}

func addItem(item *models.Item) error {
	db := getDbConnection()
	defer db.Close()

	err := db.Insert(item)
	return err
}

func deleteItem(id int64) error {
	db := getDbConnection()
	defer db.Close()

	item := &models.Item{ID: id}
	err := db.Delete(item)
	return err
}

func allItems(since int64, limit int64) (result []*models.Item) {
	db := getDbConnection()
	defer db.Close()

	var items []*models.Item
	err := db.Model(&items).Select()
	if err != nil {
		logDBIssue(err)
	}

	return items
}

// ItemService handles all the things that have to do with items
type ItemService struct{}

// Remove removes an item from your inventory or reduces
// the count if you have more than one
func (m *ItemService) Remove(id int64) error {
	return deleteItem(id)
}

// Add adds a new item to your inventory if it is not in
// your inventory and ups the count if you already have the
// item
func (m *ItemService) Add(item *models.Item) error {
	return addItem(item)
}

// Get gets a list of the items in your inventory
func (m *ItemService) Get(params *items.GetParams) []*models.Item {
	mergedParams := items.NewGetParams()
	mergedParams.Since = swag.Int64(0)
	if params.Since != nil {
		mergedParams.Since = params.Since
	}
	if params.Limit != nil {
		mergedParams.Limit = params.Limit
	}
	return allItems(*mergedParams.Since, *mergedParams.Limit)
}

// NewItemService cretes a new instance of the ItemService
func NewItemService() *ItemService {
	return &ItemService{}
}
