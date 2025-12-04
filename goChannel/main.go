package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 2)

	// A buffered channel lets you send values even when no receiver is ready, but ONLY until the buffer becomes full.
	// Buffered channel with capacity 2 can accept up to 2 sends without a receiver (non-blocking). Third send will block if no one is receiving.

	// Channels DO NOT close automatically. Buffered channels simply allow limited sends before blocking. Deadlock happens when a goroutine tries to send/receive but no other goroutine can make progress to unblock it.

	var wg sync.WaitGroup

	// g1 goroutine
	// This goroutine sends a value into the channel. If the channel were UNBUFFERED and no receiver existed, this send operation would BLOCK and the goroutine would not continue. (It's not "sleep mode" — it is BLOCKED waiting for a receiver.)

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Sending 1 from g1 goroutine")
		ch <- 1
	}()

	// g3 goroutine (Unbuffered example explanation)
	// If the channel is unbuffered and this receiver is removed: g1 blocks on "ch <- 1" because no receiver is ready, main goroutine blocks on wg.Wait(), no goroutine can continue → DEADLOCK

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()

	// 	fmt.Println("Receiving data from g3 goroutine")
	// 	data := <-ch
	// 	fmt.Println("Data =", data)
	// }()

	// second sender goroutine
	// This sends the 2nd value into the channel. Since the buffer size is 2, this still DOES NOT block.
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Sending 2 from second goroutine")
		ch <- 2
		fmt.Println("second goroutine ends")
	}()

	// third sender goroutine
	// This one WILL block forever because:
	// the buffer is already full (capacity = 2), no receiver exists to free space
	// Therefore this send blocks → wg.Done() is never reached → wg.Wait() in main blocks → DEADLOCK.

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Sending 3 from third goroutine")
		ch <- 3 //blocks here → deadlock
		fmt.Println("third goroutine ends")
	}()

	wg.Wait()
	fmt.Println("Main goroutine ends")
}

/*
	go channel is like a pipe: data enters from one side and exits from the other.
	goroutines use channels to safely exchange data.

	TWO TYPES OF CHANNELS:
	1. Unbuffered (synchronous)
	2. Buffered (asynchronous up to capacity)

	UNBUFFERED CHANNEL:
	- Has NO internal storage (0 slots).
	- send blocks until receive happens.
	- receive blocks until send happens.
	- Both operations must meet at the same moment.
	- Best for synchronization.
	- Guarantees ordering: send happens-before receive completes.
	- Not the same as buffered channel with capacity 1.

	BUFFERED CHANNEL:
	- Has multiple slots (capacity you define).
	- send blocks only when buffer is full.
	- receive blocks when buffer is empty.
	- Allows asynchronous sends up to buffer capacity.
	- Closing a buffered channel still allows receiving buffered values.

	Unbuffered channel analogy:
	- A live phone call: both sides must be active at the same time.

	Buffered channel analogy:
	- A voicemail: you drop your message and leave; they listen later.

	BLOCKING RULES:
	- Unbuffered send    → blocks until a receiver is ready.
	- Unbuffered receive → blocks until a sender is ready.
	- Buffered send      → blocks when buffer is full.
	- Buffered receive   → blocks when buffer is empty.
	- Channel ops form synchronization points (happens-before).

	CLOSING A CHANNEL:
	- close(ch) marks a channel as closed.
	- sends after close → panic.
	- receives after close → zero value, ok=false.
	- buffered values remain receivable even after close.

	DIRECTIONAL CHANNELS:
	chan<- int   // send-only
	<-chan int   // receive-only

	IMPORTANT:
	- nil channel: send/receive blocks forever.
	- nil channels are used to disable select cases.
	- channels do NOT close automatically.

	DEADLOCK:
	- Happens when all goroutines are blocked and no progress is possible.
	- Go runtime error: "fatal error: all goroutines are asleep - deadlock!"
*/
