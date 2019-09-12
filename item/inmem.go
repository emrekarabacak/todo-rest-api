package item

import (
	"fmt"

	"../domain"
)

// InMemoryRepository ...
type InMemoryRepository struct {
	currentID int
	items     []*domain.Item
}

// NewInMemoryRepository return new in memory repository instance
func NewInMemoryRepository() domain.Repository {
	return &InMemoryRepository{items: make([]*domain.Item, 0), currentID: 1}
}

// Create creates new to-do item and saves it to items slice
func (inmem *InMemoryRepository) Create(item *domain.Item) (*domain.Item, error) {
	item.ID = inmem.currentID
	inmem.currentID++
	inmem.items = append(inmem.items, item)
	return item, nil
}

// GetByID returns to-do item with given id
func (inmem *InMemoryRepository) GetByID(id int) (*domain.Item, error) {
	for _, item := range inmem.items {
		if item.ID == id {
			return item, nil
		}
	}
	return &domain.Item{}, fmt.Errorf("Item with id %d cannot be found ", id)
}

// GetAll returns all to-do items
func (inmem *InMemoryRepository) GetAll() []*domain.Item {
	return inmem.items
}

// Delete deletes to-do item with given id
func (inmem *InMemoryRepository) Delete(id int) error {
	for index, item := range inmem.items {
		if item.ID == id {
			inmem.items = append(inmem.items[:index], inmem.items[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Item with id %d cannot be found", id)
}

// Update updates to-do item with given id
func (inmem *InMemoryRepository) Update(updatedItem *domain.Item) (*domain.Item, error) {
	for index, item := range inmem.items {
		if item.ID == updatedItem.ID {
			inmem.items[index] = updatedItem
			return updatedItem, nil
		}
	}
	return &domain.Item{}, fmt.Errorf("Item with id %d cannot be found. Update failed", updatedItem.ID)
}
