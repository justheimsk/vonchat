package registry

type Registry[K string, V any] struct {
	items map[K]V
}

func NewRegistry[K string, V any]() *Registry[K, V] {
	return &Registry[K, V]{
		items: make(map[K]V),
	}
}

func (self *Registry[K, V]) Register(key K, value V) {
	self.items[key] = value
}

func (self *Registry[K, V]) Get(key K) (V, bool) {
	item, ok := self.items[key]
	return item, ok
}
