package main

import (
	"time"
)

type Item struct {
	Content      string
	CreationDate time.Time
	Completed    bool
}
