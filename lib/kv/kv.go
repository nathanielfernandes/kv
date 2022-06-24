package kv

import (
	"time"

	"github.com/cornelk/hashmap"
)

type KV struct {
	hm *hashmap.HashMap

	last_gc time.Time
	gc      []string
	next_gc []string
}

var GC_TICK = 5 * time.Hour

func NewKV() *KV {
	m := &KV{
		hm: &hashmap.HashMap{},

		last_gc: time.Now(),
		gc:      []string{},
		next_gc: []string{},
	}
	go m.start()
	return m
}

func (m *KV) Get(key string, def string) string {
	if v, ok := m.hm.Get(key); ok {
		return v.(string)
	}
	return def
}

func (m *KV) Set(key string, value string) {
	m.hm.Set(key, value)
	m.next_gc = append(m.next_gc, key)
}

func (m *KV) clean() {
	for _, key := range m.gc {
		m.hm.Del(key)
	}
	m.gc = m.next_gc
	m.next_gc = []string{}
}

func (m *KV) start() {
	for range time.Tick(GC_TICK) {
		m.clean()
		m.last_gc = time.Now()
	}
}
