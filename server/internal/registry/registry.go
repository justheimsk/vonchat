package registry

import (
	"sync"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

type Registry[K comparable, V any] struct {
	items map[K]V
	mu    sync.Mutex
}

func NewRegistry[K comparable, V any]() *Registry[K, V] {
	return &Registry[K, V]{
		items: make(map[K]V),
	}
}

func (self *Registry[K, V]) Register(key K, value V) error {
	_, found := self.Get(key)
	if found {
		return models.ErrDuplicate
	}

	self.mu.Lock()
	defer self.mu.Unlock()

	self.items[key] = value
	return nil
}

func (self *Registry[K, V]) Get(key K) (V, bool) {
	self.mu.Lock()
	defer self.mu.Unlock()

	item, ok := self.items[key]
	return item, ok
}

func (self *Registry[K, V]) Remove(key K) {
	self.mu.Lock()
	defer self.mu.Unlock()

	delete(self.items, key)
}

func (self *Registry[K, V]) Values() map[K]V {
	return self.items
}
