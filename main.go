package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func EncodeJSON() []byte {
	// a sample json
	a := map[string]string{
		"field_1": "this is one field",
		"field_2": "this is second field",
		"field_3": "this is third field",
	}

	encoded, err := json.Marshal(a)
	if err != nil {
		log.Fatalf("Error encoding JOSN data - %v", err)
	}
	return encoded
}

func main() {
	fmt.Println("Starting .........")
	fmt.Printf("simple use of memcached in golang\n\n")

	// Connect to store and add value to it
	my_value := &memcache.Item{
		Key:        "foo",
		Value:      EncodeJSON(),
		Expiration: 10,
	}

	mc := memcache.New("127.0.0.1:11211")

	item, err := mc.Get("foo")

	// check if any error or cache miss
	if err != nil {
		if err == memcache.ErrCacheMiss {
			log.Println("Cache miss!!")
			log.Printf("Adding new value with key: %v\n", my_value.Key)
			mc.Set(my_value)
		} else {
			log.Fatalf("Error fetching from memcache - %v", err)
		}
	}

	// check again
	item, _ = mc.Get("foo")

	fmt.Printf("Item Value: %s\n", (item.Value))
}
