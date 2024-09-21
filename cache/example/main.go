package main

import (
	"fmt"
	"time"

	gocache "github.com/rahulshewale153/common-lib/cache/gocache"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// initialise the cache of gocache and redis
	gocache.NewGoCache(1*time.Second, 1*time.Second)

	gocache.Set("foo", "bar", 2*time.Second)
	value, found := gocache.Get("foo")
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Key not found")
	}

}
