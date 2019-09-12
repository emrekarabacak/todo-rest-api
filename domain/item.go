package domain

import "time"

// Item represents single to-do item
type Item struct {
	ID        int       `json:"id"`
	Content   string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Completed bool      `json:"completed"`
}

// Service for managing to-do items
type Service interface {
	CreateItem(content string) (*Item, error)
	FindItemByID(id int) (*Item, error)
	FindAll() []*Item
	DeleteByID(id int) error
	CompleteItem(id int) (*Item, error)
}

// Repository CRUD operations on to-do items
type Repository interface {
	Create(item *Item) (*Item, error)
	GetByID(id int) (*Item, error)
	GetAll() []*Item
	Delete(id int) error
	Update(item *Item) (*Item, error)
}

// ItemRequest represent new item request body
type ItemRequest struct {
	Content string `json:"content"`
}
