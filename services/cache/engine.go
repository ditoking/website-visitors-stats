package cache

import (
	"github.com/patrickmn/go-cache"
)

type Engine struct {
	cacheSrv *cache.Cache
}

func NewEngine(cacheSrv *cache.Cache) *Engine {
	e := new(Engine)
	e.cacheSrv = cacheSrv

	return e
}

func (e Engine) CreateBoolEntryIfNotExists(key string) (bool, error) {
	_, found := e.cacheSrv.Get(key)

	if !found {
		e.cacheSrv.SetDefault(key, true)
	}

	return found, nil
}

func (e Engine) IncrementByKey(key string, incrementValue int) error {
	_, found := e.cacheSrv.Get(key)

	if !found {
		e.cacheSrv.SetDefault(key, int64(incrementValue))
		return nil
	} else {
		return e.cacheSrv.Increment(key, int64(incrementValue))
	}
}

func (e Engine) GetIntValueByKey(key string) int64 {
	value, found := e.cacheSrv.Get(key)

	if !found {
		return 0
	}

	return value.(int64)
}