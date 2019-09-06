package main

import (
	"errors"
	"time"
)

var items = make([]Item, 0)
var currentId = 1

func GetAll() []Item {
	return items
}

func Add(content string) Item {
	item := Item{Id: currentId, Content: content, CreationDate: time.Now(), Completed: false}
	items = append(items, item)
	currentId++
	return item
}

func Get(id int) (Item, error) {
	for _, item := range items {
		if item.Id == id {
			return item, nil
		}
	}
	return Item{}, errors.New("Item cannot be found")
}

func Update(id int, completeStatus bool) (Item, error) {
	for index, item := range items {
		if item.Id == id {
			items[index].Completed = completeStatus
			return items[index], nil
		}
	}
	return Item{}, errors.New("Item cannot be found")
}

func Delete(id int) error {
	for index, item := range items {
		if item.Id == id {
			items = append(items[:index], items[index+1:]...)
			return nil
		}
	}
	return errors.New("Item cannot be found")
}
