package cache

import "sync"

type CacheData struct {
	usersStatus map[string]string
	mu          sync.RWMutex
}

type InMemoryCache struct {
	data CacheData
}

var (
	instance *InMemoryCache
	once     sync.Once
)

func NewInMemoryCache() *InMemoryCache {
	once.Do(func() {
		instance = &InMemoryCache{
			data: CacheData{
				usersStatus: make(map[string]string),
			},
		}
	})

	return instance
}

func (self *InMemoryCache) GetUserStatus(id string) string {
	self.data.mu.Lock()
	defer self.data.mu.Unlock()

	status, found := self.data.usersStatus[id]
	if !found {
		return ""
	}

	return status
}

func (self *InMemoryCache) SetUserStatus(id string, status string) {
	self.data.mu.Lock()
	defer self.data.mu.Unlock()

	self.data.usersStatus[id] = status
}
