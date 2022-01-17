package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Proxy interface {
	Accept(key string) bool
	Proxy(c *fiber.Ctx) error
}

// proxy1

// proxy2

var Proxies = []Proxy{

	NewLimitProxy("user", 3, 3*time.Minute),
	// NewCacheProxy("user", 3*time.Minute),
	// NewCacheProxy("event", 10*time.Minute),
}

func ProxyHandler(c *fiber.Ctx) error {
	for _, v := range Proxies {
		if v.Accept(c.Params("key")) {
			return v.Proxy(c)
		}
	}
	c.Response().SetStatusCode(404)
	return nil
}
