package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex // guards mapping
	mapping    map[string]string
}{
	mapping: make(map[string]string),
}

func main() {
	//!+main
	cache.Lock()
	cache.mapping["bob"] = "jones"
	cache.mapping["joe"] = "smith"
	cache.Unlock()

	fmt.Printf("%v,%v\n", Lookup("bob"), cache.test)

	//!-main
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
