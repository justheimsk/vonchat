package registry

import "github.com/justheimsk/vonchat/server/internal/domain/models"

type Registry[K comparable, V any] struct {
	items map[K]V
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

	self.items[key] = value
	return nil
}

func (self *Registry[K, V]) Get(key K) (V, bool) {
	item, ok := self.items[key]
	return item, ok
}

func (self *Registry[K, V]) Remove(key K) {
	delete(self.items, key)
}

func (self *Registry[K, V]) Values() map[K]V {
	return self.items
}
