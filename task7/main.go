package main

import (
	"fmt"
	"strings"
	"sync"
)

// Map is a concurrent-safe map wrapper, which key
//is of any comparable type, and value is just of any type
type Map[K comparable, V any] struct {
	m  map[K]V     // underlying map with generic key and value types
	mu *sync.Mutex // underlying mutex for concurrency-safety
}

// New is a constructor function for Map type
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:  make(map[K]V),
		mu: &sync.Mutex{},
	}
}

// Set method stores the key-value pair within the underlying map
func (m *Map[K, V]) Set(key K, val V) {
	m.mu.Lock()    // mutex ON
	m.m[key] = val // set the key-value pair
	m.mu.Unlock()  // mutex OFF
}

// Get method retrieves value associated with a given key
func (m Map[K, V]) Get(key K) V {
	m.mu.Lock()     // mutex ON
	res := m.m[key] // get the value
	m.mu.Unlock()   // mutex OFF
	return res
}

// GetOK method retrieves value associated with a given key,
// and true if value exists, or zero-value and false if not
func (m Map[K, V]) GetOK(key K) V {
	m.mu.Lock()     // mutex ON
	res := m.m[key] // get the value
	m.mu.Unlock()   // mutex OFF
	return res
}

// Delete method removes a key-value pair from
// an underlying map
func (m *Map[K, V]) Delete(key K) {
	m.mu.Lock()      // mutex ON
	delete(m.m, key) // delete a key-value pair
	m.mu.Unlock()    // mutex OFF
}

// Has method returns either true if Map has a value,
//associated with a given key or false if not
func (m Map[K, V]) Has(key K) bool {
	m.mu.Lock()       // mutex ON
	_, ok := m.m[key] // get bool
	m.mu.Unlock()     // mutex OFF
	return ok
}

func (m Map[K, V]) String() string {
	sb := strings.Builder{} // init strings.Builder
	sb.WriteString("{ ")
	for k, v := range m.m {
		sb.WriteString(fmt.Sprintf("%v => %v, ", k, v))
	}
	sb.WriteString(" }")
	return sb.String() // return accumulated string
}

func main() {
	m := New[int, string]()            // init Map
	a := []string{"foo", "bar", "baz"} // init data
	wg := sync.WaitGroup{}             // init sync.WaitGroup
	wg.Add(len(a))                     // inc WG counter

	// concurrent setting
	for i, val := range a {
		go func(j int, v string) {
			defer wg.Done()
			m.Set(j, v)
		}(i, val)
	}

	wg.Wait()      // waiting for goroutines to finish their jobs
	wg.Add(len(a)) // inc WG counter once again

	// concurrent reading
	for i := range a {
		go func(j int) {
			defer wg.Done()
			fmt.Println(m.Get(j))
		}(i)
	}

	wg.Wait()      // waiting for goroutines to finish their jobs
	wg.Add(len(a)) // inc WG counter once again

	fmt.Println(m)

	// concurrent deleting
	for i := range a {
		go func(j int) {
			defer wg.Done()
			m.Delete(j)
		}(i)
	}

	wg.Wait()
	fmt.Println(m)
}
