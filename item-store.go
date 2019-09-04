package main

import "time"

var items = make([]Item, 0)

func GetAll() []Item {
	return items
}

func Add(content string) Item {
	item := Item{Content: content, CreationDate: time.Now(), Completed: false}
	items = append(items, item)
	return item
}
