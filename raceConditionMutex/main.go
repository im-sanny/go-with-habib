package main

import (
	"fmt"
	"sync"
)

var counter int64
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {

			// Rule: Lock and unlock in the same goroutine.
			// Rule: Keep critical sections small.
			mu.Lock()
			// Rule: Use defer for unlock unless in hot path (omitted here for clarity).
			defer mu.Unlock()

			// Rule: Never mutate shared state without holding the lock.
			// (If we removed mu.Lock(), this line would cause a race.)
			a := counter
			a = a + 1 // Race occurs here if mutex is removed.
			counter = a

			wg.Done()
		}()

	}

	// Rule: Do not start goroutines while holding a lock
	wg.Wait()

	fmt.Println(counter)
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


	-	Mutex:
		A mutex (mutual exclusion lock) is a synchronization mechanism that ensures only one goroutine can access a shared resource at a time.
		It prevents race conditions by allowing exclusive access—when one goroutine holds the lock, all others must wait.

	- Imagine a file as a room and the mutex as a lock. When Person A enters, they lock the door. B and C must wait outside. A safely updates the file, finishes, unlocks the door, and leaves. Then B or C can enter. This lock–unlock system prevents two people from modifying the file at the same time, avoiding race conditions.

	 Mutex Rules:
	 - Never copy a mutex or a struct containing one.
	 - Lock and unlock in the same goroutine.
	 - Keep critical sections small; avoid I/O/blocking while locked.
	 - Use `defer mu.Unlock()` right after `mu.Lock()` (except hot paths).
	 - Maintain a global lock acquisition order (prevents deadlocks).
	 - Do not expose mutexes via public APIs.
	 - Under RLock: never mutate shared state.
	 - Do not upgrade RLock -> Lock; release, then Lock().
	 - Never start goroutines while holding a lock.
	 - Avoid channel send/recv inside locked regions.
	 - Use RWMutex only when read contention is proven high.

*/
