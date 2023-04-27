package demo

import (
	"sync"
)

type MyCache struct {
	pool sync.Pool
}

func NewMyCache() *MyCache {
	return &MyCache{}
}
