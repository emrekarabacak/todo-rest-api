package main

import (
	"time"
)

type Item struct {
	Id           int
	Content      string
	CreationDate time.Time
	Completed    bool
}
