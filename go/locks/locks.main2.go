package main

import (
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/diaevd/lock"
)

func main() {
	kl := lock.NewKeyLock()
	lerr := log.New(os.Stderr, "", 0xff)
	lerr.Printf("Run test\n")
	var wg sync.WaitGroup

	wg.Add(2)
	go func(sleep int) {
		defer wg.Done()
		lerr.Printf("func(%v):0 s: %v\n", sleep, time.Second*time.Duration(sleep))
		kl.Block("key1")
		lerr.Printf("func(%v):1 s: %v\n", sleep, time.Second)
		time.Sleep(time.Second * time.Duration(sleep))
		lerr.Printf("func(%v):2 s: %v\n", sleep, time.Second)
		kl.Unblock("key1")
		lerr.Printf("func(%v):3 s: %v\n", sleep, time.Second)
		runtime.Gosched() // на всяк случай
	}(5)

	go func(sleep int) {
		defer wg.Done()
		lerr.Printf("func(%v):0 s: %v\n", sleep, time.Second*time.Duration(sleep))
		kl.Block("key1")
		lerr.Printf("func(%v):1 s: %v\n", sleep, time.Second)
		time.Sleep(time.Second * time.Duration(sleep))
		lerr.Printf("func(%v):2 s: %v\n", sleep, time.Second)
		kl.Unblock("key2")
		lerr.Printf("func(%v):3 s: %v\n", sleep, time.Second)
		runtime.Gosched() // на всяк случай
	}(5)

	// time.Sleep(time.Second * 15)
	wg.Wait()

	lerr.Printf("End test")
}
