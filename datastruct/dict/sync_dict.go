package dict

import "sync"

type SyncDict struct {
	m sync.Map
}

func (dict *SyncDict) Get(key string) (val interface{}, exists bool) {
	value, ok := dict.m.Load(key)
	return value, ok
}

func (dict *SyncDict) Len() int {
	lenth := 0
	dict.m.Range(func(key, value interface{}) bool {
		lenth++
		return true
	})
	return lenth
}

func (dict *SyncDict) Put(key string, val interface{}) (result int) {
	_, existed := dict.m.Load(key)
	dict.m.Store(key, val)
	if existed {
		return 0
	}
	return 1
}

func (dict *SyncDict) PutIfAbsent(key string, val interface{}) (result int) {
	_, existed := dict.m.Load(key)
	if existed {
		return 0
	}
	dict.m.Store(key, val)
	return 1
}
func (dict *SyncDict) PutIfExists(key string, val interface{}) (result int) {
	_, existed := dict.m.Load(key)
	if !existed {
		return 0
	}
	dict.m.Store(key, val)
	return 1
}
func (dict *SyncDict) Remove(key string) (result int) {
	_, existed := dict.m.Load(key)
	dict.m.Delete(key)
	if existed {
		return 1
	}
	return 0
}

func (dict *SyncDict) ForEach(consumer Consumer) {
	dict.m.Range(consumer)
}

func (dict *SyncDict) Keys() []string {
	//TODO implement me
	panic("implement me")
}

func (dict *SyncDict) RandomKeys(limit int) []string {
	//TODO implement me
	panic("implement me")
}

func (dict *SyncDict) RandomDistinctKeys(limit int) []string {
	//TODO implement me
	panic("implement me")
}

func (dict *SyncDict) clear() {
	//TODO implement me
	panic("implement me")
}

func MakeSyncDict() *SyncDict {
	return &SyncDict{}
}
