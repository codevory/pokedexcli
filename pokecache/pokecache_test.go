package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestPokeGetCache(t *testing.T) {
	cache := NewCache(time.Second * 5)
	cache.Add("cache1", []byte("pikachu"))

	getCacheCases := []struct {
		key          string
		expectedByte []byte
		expectedOk   bool
	}{
		{
			key:          "cache1",
			expectedByte: []byte("pikachu"),
			expectedOk:   true,
		},
		{
			key:          "cache2",
			expectedByte: nil,
			expectedOk:   false,
		},
		{
			key:          "cache3",
			expectedByte: nil,
			expectedOk:   false,
		},
	}

	for _, gc := range getCacheCases {
		bytesGot, ok := cache.Get(gc.key)
		if ok != gc.expectedOk {
			t.Errorf("For key %q:, expected ok = %v, got : %v", gc.key, gc.expectedOk, ok)
		}

		if !bytes.Equal(bytesGot, gc.expectedByte) {
			t.Errorf("for key %q: expected bytes=%s, got %s", gc.key, string(gc.expectedByte), string(bytesGot))
		}
	}
}

func TestPokeAddCache(t *testing.T) {
	cache := NewCache(time.Second * 5)

	addCacheCases := []struct {
		key      string
		val      []byte
		expected bool
	}{
		{
			key:      "cache1",
			val:      []byte("pikachu"),
			expected: true,
		},
		{
			key:      "cache2",
			val:      []byte("hello! world"),
			expected: true,
		},
		{
			key:      "",
			val:      []byte("hello baby"),
			expected: false,
		},
		{
			key:      "cache1",
			val:      nil,
			expected: false,
		},
	}

	for _, ac := range addCacheCases {
		got := cache.Add(ac.key, ac.val)
		if got != ac.expected {
			t.Errorf("expected : %v, got %v", ac.expected, got)
		}
	}
}
