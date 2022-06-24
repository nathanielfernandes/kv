package kv

import (
	"time"

	"github.com/cornelk/hashmap"
)

type KV struct {
	hm *hashmap.HashMap

	last_gc time.Time
	gc      *hashmap.HashMap
	next_gc *hashmap.HashMap
}

var GC_TICK = 6 * time.Hour

func NewKV() *KV {
	m := &KV{
		hm: &hashmap.HashMap{},

		last_gc: time.Now(),
		gc:      &hashmap.HashMap{},
		next_gc: &hashmap.HashMap{},
	}
	go m.start()
	return m
}

func (m *KV) reset_exipirey(key string) {
	if _, ok := m.gc.Get(key); ok {
		m.gc.Del(key)
		m.next_gc.Set(key, struct{}{})
	}
}

func (m *KV) Get(key string, def string) string {
	if v, ok := m.hm.GetStringKey(key); ok {
		m.reset_exipirey(key)
		return v.(string)
	}
	return def
}

func (m *KV) Set(key string, value string) bool {
	if _, ok := m.hm.Get(key); ok {
		return false
	}

	m.hm.Set(key, value)
	m.next_gc.Set(key, struct{}{})
	return true
}

func (m *KV) clean() {
	for kv := range m.gc.Iter() {
		m.hm.Del(kv.Key)
	}
	m.gc = m.next_gc
	m.next_gc = &hashmap.HashMap{}
}

func (m *KV) start() {
	for range time.Tick(GC_TICK) {
		m.clean()
		m.last_gc = time.Now()
	}
}
