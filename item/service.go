package item

import (
	"time"
	"todo-rest-api/domain"
)

// Service to-do items business logic
type Service struct {
	repository domain.Repository
}

// NewItemService create new service instance
func NewItemService(repository domain.Repository) domain.Service {
	return &Service{repository: repository}
}

// CreateItem creates new to-do item
func (itemService *Service) CreateItem(content string) (*domain.Item, error) {
	return itemService.repository.Create(&domain.Item{Content: content, CreatedAt: time.Now(), Completed: false})
}

// FindItemByID returns to-do item with given id
func (itemService *Service) FindItemByID(id int) (*domain.Item, error) {
	return itemService.repository.GetByID(id)
}

// FindAll return all to-do items
func (itemService *Service) FindAll() []*domain.Item {
	return itemService.repository.GetAll()
}

// DeleteByID deletes to-do item with given id
func (itemService *Service) DeleteByID(id int) error {
	return itemService.repository.Delete(id)
}

// CompleteItem sets to-do item as completed
func (itemService *Service) CompleteItem(id int) (*domain.Item, error) {
	item, err := itemService.repository.GetByID(id)
	if err != nil {
		return &domain.Item{}, err
	}
	item.Completed = true
	return itemService.repository.Update(item)
}
