package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// We plan to start 5 goroutines
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		// Capture the value of i
		i := i

		go func() {
			defer wg.Done() // Same as wg.Add(-1) when this goroutine finishes
			fmt.Printf("Worker %d starting\n", i)
			time.Sleep(time.Duration(i) * 200 * time.Millisecond)
			fmt.Printf("Worker %d done\n", i)
		}()
	}

	fmt.Println("Main waiting for all workers...")
	wg.Wait() // Blocks until counter becomes 0
	fmt.Println("All workers finished. Main exits.")
}

/*

wg := WaitGroup{
		noCopy: noCopy{}
		state: atomic.Uint64{
			_ : noCopy{}
			_ : align64{}
			v: 0
			}
			sema:0,
		}

- Waitgroup:
		If you call Done() without calling Add() first, the WaitGroup will panic and crash the program.

		If the WaitGroup counter is set to 5 but only three Done() calls happen, the counter never reaches zero, so any goroutine waiting on it will block indefinitely.

		64 bit of states uint64 are divided in 2 part, they are high 32 bit and low 32 bit. high 32 bit also know as counter or unfinished worker counter/unfinished goroutine, this high32 bit keeps those who's yet to run or unfinished. after executing all the goroutine, say we had 3 programs in high32 and go routine executed them asynchronously, and when these are done the main go routine will call the wg.wait() in the last and it'll go check the high32 and if it sees the number is 0 then it won't stay stuck anymore and it'll move to next line of the code(if the high32 bit =0, this goroutine will not sleep)

- rightshift operator:


- deadlock: I am expecting some to give me something and waiting for that but that person will never give me the stuff and my waiting for that stuff is deadlock.

	A sync.WaitGroup is used to wait for a collection of goroutines to finish. No API changes in Go 1.25.

- Three main methods:
	- wg.Add(delta int)     → increase the counter (usually call before go func())
	- wg.Done()             → decrease the counter by 1 (call just before goroutine ends)
	- wg.Wait()             → blocks current goroutine until counter == 0

- Important rules:
	1. Add() must be called BEFORE starting the goroutine (otherwise race!)
	2. Every Add(1) must have exactly one Done()
	3. You can call Add(n) once for n goroutines instead of calling Add(1) n times
	4. Done() is exactly the same as Add(-1)
	5. Calling Add() with positive value after Wait() returns → panic
	6. Calling Done() too many times → panic

*/
