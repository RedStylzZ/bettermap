package bettermap

import "sync"

type BetterMap[K comparable, V any] struct {
	sync.Mutex
	m    map[K]V
	keys []K
}

func NewBetterMap[K comparable, V any]() *BetterMap[K, V] {
	m := make(map[K]V)
	return &BetterMap[K, V]{m: m}
}

func (b *BetterMap[K, V]) Get(key K) V {
	v, _ := b.GetAndCheck(key)
	return v
}

func (b *BetterMap[K, V]) GetAndCheck(key K) (V, bool) {
	v, ok := b.m[key]
	return v, ok
}

func (b *BetterMap[K, V]) GetMany(keys []K) []V {
	values := make([]V, 0)
	for _, key := range keys {
		values = append(values, b.m[key])
	}
	return values
}

func (b *BetterMap[K, V]) GetByValue(f func(value V) bool) []V {
	values := make([]V, 0)
	for _, key := range b.keys {
		value := b.m[key]
		if f(value) {
			values = append(values, value)
		}
	}
	return values
}

func (b *BetterMap[K, V]) Set(key K, value V) (V, bool) {
	b.Lock()
	defer b.Unlock()

	b.addKeyToList(key)
	b.m[key] = value
	v, ok := b.m[key]
	return v, ok
}

func (b *BetterMap[K, V]) Remove(key K) {
	b.Lock()
	defer b.Unlock()

	delete(b.m, key)

	if idx := b.keyIndex(key); idx != -1 {
		b.keys = append(b.keys[:idx], b.keys[idx+1:]...)
	}
}

func (b *BetterMap[K, V]) Keys() []K {
	return b.keys
}

func (b *BetterMap[K, V]) Values() []V {
	values := make([]V, len(b.m))

	for i, key := range b.keys {
		values[i] = b.m[key]
	}

	return values
}

func (b *BetterMap[K, V]) Raw() map[K]V {
	return b.m
}

func (b *BetterMap[K, V]) keyIndex(key K) int {
	for i, lKey := range b.keys {
		if lKey == key {
			return i
		}
	}
	return -1
}

func (b *BetterMap[K, V]) addKeyToList(key K) {
	if _, exist := b.m[key]; exist {
		return
	}
	b.keys = append(b.keys, key)
}
