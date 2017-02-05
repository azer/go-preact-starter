package preact

type Cache struct {
	Content map[string]string
}

func (cache *Cache) Get(path string) (string, bool) {
	if _, ok := cache.Content[path]; !ok {
		return "", false
	}

	return cache.Content[path], true
}

func (cache *Cache) Set(path, html string) {
	cache.Content[path] = html
}

func NewCache() *Cache {
	return &Cache{
		Content: map[string]string{},
	}
}
