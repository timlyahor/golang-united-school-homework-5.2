package cache

import "time"

type Cache struct {
	holder map[string]Object
}

type Object struct {
	value    string
	isTill   bool
	deadline time.Time
}

func NewCache() Cache {
	return Cache{
		holder: make(map[string]Object),
	}
}

func (cache *Cache) Get(key string) (string, bool) {
	value, isExist := cache.holder[key]
	if isExist == false {
		return "", isExist

	} else {
		if value.isTill == false {
			return value.value, true
		} else {
			now := time.Now()

			if now.After(value.deadline) || now.Equal(value.deadline) {
				delete(cache.holder, key)
				return "", false
			}

			return value.value, true
		}
	}
}

func (cache *Cache) Put(key, value string) {
	cache.holder[key] = Object{
		value: value,
	}
}

func (cache *Cache) Keys() []string {
	res := make([]string, 0)
	for k, v := range cache.holder {
		if v.isTill == false {
			res = append(res, k)
		} else {
			now := time.Now()

			if now.Before(v.deadline) {
				res = append(res, k)
			}
		}
	}

	return res
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.holder[key] = Object{
		value:    value,
		isTill:   true,
		deadline: deadline,
	}
}
