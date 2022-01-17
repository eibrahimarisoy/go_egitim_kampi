package main

import (
	"time"
)

var counter = map[string]*Cache{}

type Cache struct {
	body []byte
	ttl  time.Time
}

type CacheProxy struct {
	key string
	ttl time.Duration
}
