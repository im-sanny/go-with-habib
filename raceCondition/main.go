package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter = counter + 1 // <- RACE HAPPENS HERE
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter =", counter)
}

/*
- Race Condition:
	A race condition occurs when two or more goroutines access the same shared variable concurrently and at least one access is a write, and the final result depends on the unpredictable execution order.
	Three required conditions (all must be true):

	- Shared data – a variable visible to multiple goroutines
	- Concurrent access – multiple goroutines running at the same time
	- No proper synchronization – no mutex, channel, atomic ops, etc.

	Typical symptoms:
	- Non-deterministic output (different values on each run)
	- Wrong/incorrect final values (like counter < 1000 above)
	- Hard-to-reproduce bugs

	How to Fix Race Conditions

	Use any synchronization tool that ensures goroutines don’t modify data simultaneously:
	Mutex:
	- Prevents multiple goroutines from entering the critical section.
	Atomic Operations:
	- Lock-free and very fast for simple counters.
	Channels:
	- Go’s recommended method — allows safe communication instead of sharing memory.

	Go philosophy:
	- Do not communicate by sharing memory; instead, share memory by communicating.

	go run -race main.go    # Use this command to detect race condition, it's the built-in race detector
*/
