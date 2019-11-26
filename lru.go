package lru

type LRU struct {
	size     int
	cache    map[interface{}]interface{}
	oldCache map[interface{}]interface{}
}

func New(size int) *LRU {
	l := &LRU{
		size: size,
	}
	l.Clear()
	return l
}

func (l *LRU) Set(key, value interface{}) {
	l.cache[key] = value

	if len(l.cache) >= l.size {
		l.oldCache = l.cache
		l.cache = map[interface{}]interface{}{}
	}
}

func (l *LRU) Get(key interface{}) (value interface{}, ok bool) {
	if value, ok = l.cache[key]; ok {
		return
	}

	if value, ok = l.oldCache[key]; ok {
		delete(l.oldCache, key)
		l.Set(key, value)
		return
	}

	return
}

func (l *LRU) Peek(key interface{}) (value interface{}, ok bool) {
	if value, ok = l.cache[key]; ok {
		return
	}

	if value, ok = l.oldCache[key]; ok {
		return
	}

	return
}

func (l *LRU) Has(key interface{}) bool {
	if _, ok := l.cache[key]; ok {
		return true
	}

	if _, ok := l.oldCache[key]; ok {
		return true
	}

	return false
}

func (l *LRU) Remove(key interface{}) {
	delete(l.cache, key)
	delete(l.oldCache, key)
}

func (l *LRU) Clear() {
	l.cache = map[interface{}]interface{}{}
	l.oldCache = map[interface{}]interface{}{}
}
