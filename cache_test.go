/*
Package timeCache is a simple implementation of an interface and two structures
which define aa in-memory cache where entries have a timeout


Released under MIT license, copyright 2018 Tyler Ramer
*/

package timeCache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const TIMEOUT = 30
const CacheSize = 10000000 // test includes key 1,000,000 so set to higher than this

var timeout = time.Duration(TIMEOUT) * time.Second

var s SliceCache
var d DictCache

func TestSliceAdd(t *testing.T) {
	s.timeout = timeout

	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		s.Push(key)
	}
}

func TestSliceContains1(t *testing.T) {
	var c bool
	c, _ = s.Contains("key-1")
	if !c {
		t.Log("key-1 should be included in this cache")
		t.Fail()
	}
}

func TestSliceContains1000000(t *testing.T) {
	var c bool
	c, _ = s.Contains("key-1000000")
	if !c {
		t.Log("key-1000000 should be included in this cache")
		t.Fail()
	}
}

func TestSliceDoesNotContain(t *testing.T) {
	var c bool
	size := s.count
	t.Logf("Current size %d", size)

	c, _ = s.Contains("keyTEST")

	if c {
		t.Log("keyTEST should NOT be included in this cache yet")
		t.Fail()
	}
	if s.count != size+1 {
		t.Logf("Cache size should have grown by 1; current size: %d", s.count)
		t.Fail()
	}
	c, _ = s.Contains("keyTEST")

	if !c {
		t.Log("keyTEST has been added and should now be in the cache")
		t.Fail()
	}
}

func TestSliceExpire(t *testing.T) {
	time.Sleep(timeout)

	c, _ := s.Contains("key-1")

	if s.count != 1 {
		t.Log("Size of cache is not 1 , but all entries should be expired, then one added")
		t.Fail()
	}
	if c {
		t.Log("key should have expired")
		t.Fail()
	}
}

func TestDictCreate(t *testing.T) {
	testDictCache := NewDictCache()
	if testDictCache.count > 0 {
		t.Log("new cache should have 0 entires")
		t.Fail()
	}
	c, _ := testDictCache.Contains("TEST")
	if c {
		t.Log("New dict cache should not contain a value")
		t.Fail()
	}
	if testDictCache.count != 1 {
		t.Log("Test cache should only contain one entry")
		t.Fail()
	}
}

func TestDictAdd(t *testing.T) {
	d.timeout = timeout
	d.entries = make(map[string]time.Time)

	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		d.Push(key)
	}
}

func TestDictContains1(t *testing.T) {
	var c bool
	c, _ = d.Contains("key-1")
	if !c {
		t.Log("key-1 should be included in this cache")
		t.Fail()
	}
}

func TestDictContains1000000(t *testing.T) {
	var c bool
	c, _ = d.Contains("key-1000000")
	if !c {
		t.Log("key-1000000 should be included in this cache")
		t.Fail()
	}
}

func TestDictDoesNotContain(t *testing.T) {
	var c bool
	size := d.count
	t.Logf("Current size %d", size)

	c, _ = d.Contains("keyTEST")
	if c {
		t.Log("keyTEST should NOT be included in this cache yet")
		t.Fail()
	}
	if d.count != size+1 {
		t.Logf("Cache size should have grown by 1; current size: %d", d.count)
		t.Fail()
	}
	c, _ = d.Contains("keyTEST")

	if !c {
		t.Log("keyTEST has been added and should now be in the cache")
		t.Fail()
	}
}

func TestDictExpire(t *testing.T) {
	time.Sleep(timeout)

	c, _ := d.Contains("key-1")
	if d.count != 1 {
		t.Log("Size of cache is not 1 , but all entries should be expired, then one added")
		t.Fail()
	}

	if c {
		t.Log("key should have expired")
		t.Fail()
	}
}

func BenchmarkSliceDoesNotContainRand(b *testing.B) {
	var sBench SliceCache
	sBench.timeout = timeout
	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		sBench.Push(key)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("NONkey-%d", rand.Intn(CacheSize-1))
		sBench.Contains(key)
	}
}

func BenchmarkSliceContainsRand(b *testing.B) {
	var sBench SliceCache
	sBench.timeout = timeout
	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		sBench.Push(key)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", rand.Intn(CacheSize-1))
		sBench.Contains(key)
	}
}

func BenchmarkSliceContainsLow(b *testing.B) {
	var sBench SliceCache
	sBench.timeout = timeout
	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		sBench.Push(key)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", rand.Intn(100))
		sBench.Contains(key)
	}
}

func BenchmarkSliceContainsHigh(b *testing.B) {
	var sBench SliceCache
	sBench.timeout = timeout
	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		sBench.Push(key)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", rand.Intn(100)+CacheSize-101)
		sBench.Contains(key)
	}
}

func BenchmarkDictDoesNotContainRand(b *testing.B) {
	var dBench DictCache
	dBench.timeout = timeout
	dBench.entries = make(map[string]time.Time)
	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		dBench.Push(key)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("NONkey-%d", rand.Intn(CacheSize-1))
		dBench.Contains(key)
	}
}

func BenchmarkDictContainsRand(b *testing.B) {
	var dBench DictCache
	dBench.timeout = timeout
	dBench.entries = make(map[string]time.Time)
	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		dBench.Push(key)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", rand.Intn(CacheSize-1))
		dBench.Contains(key)
	}
}

func BenchmarkDictContainsLow(b *testing.B) {
	var dBench DictCache
	dBench.timeout = timeout
	dBench.entries = make(map[string]time.Time)
	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		dBench.Push(key)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", rand.Intn(100))
		dBench.Contains(key)
	}
}

func BenchmarkDictContainsHigh(b *testing.B) {
	var dBench DictCache
	dBench.timeout = timeout
	dBench.entries = make(map[string]time.Time)
	for i := 0; i < CacheSize; i++ {
		key := fmt.Sprintf("key-%d", i)
		dBench.Push(key)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", rand.Intn(100)+CacheSize-101)
		dBench.Contains(key)
	}
}
