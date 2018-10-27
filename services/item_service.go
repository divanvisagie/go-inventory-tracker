package services

import (
	"sync"
	"sync/atomic"

	"github.com/divanvisagie/go-inventory-tracker/models"
	"github.com/divanvisagie/go-inventory-tracker/restapi/operations/items"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

var itemsStore = make(map[int64]*models.Item)
var lastID int64

var itemsLock = &sync.Mutex{}

func newItemID() int64 {
	return atomic.AddInt64(&lastID, 1)
}

func addItem(item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	newID := newItemID()
	item.ID = newID
	itemsStore[newID] = item

	return nil
}

func deleteItem(id int64) error {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := itemsStore[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(itemsStore, id)
	return nil
}

func allItems(since int64, limit int64) (result []*models.Item) {
	result = make([]*models.Item, 0)
	for id, item := range itemsStore {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, item)
		}
	}
	return
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
