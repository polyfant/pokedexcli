package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestCacheAdd(t *testing.T) {
	cache := NewCache(time.Minute)
	
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "test1",
			value: []byte("value1"),
		},
		{
			key:   "test2",
			value: []byte("value2"),
		},
	}

	for _, c := range cases {
		cache.Add(c.key, c.value)
		got, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("Expected to find key %s", c.key)
		}
		if string(got) != string(c.value) {
			t.Errorf("Expected value %s, got %s", string(c.value), string(got))
		}
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := NewCache(time.Millisecond)
	cache.Add("test", []byte("value"))
	
	time.Sleep(time.Millisecond * 2)
	
	_, ok := cache.Get("test")
	if ok {
		t.Error("Expected key to be expired")
	}
}

