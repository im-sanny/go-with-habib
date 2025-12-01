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
SYNC.WAITGROUP –
• Purpose
  WaitGroup lets the main goroutine (or any goroutine) wait until a group of goroutines finish.

 - What it does (super simple)
   WaitGroup = "wait for many goroutines to finish".
   Main goroutine says: "I pause here until all workers are done."

 - The 3 methods you use every day
   wg.Add(n)   → "I will launch n goroutines"          (call BEFORE go func())
   wg.Done()   → "one goroutine just finished"         (call at the end)
   wg.Wait()   → "sleep until all are finished"

   Done() is exactly the same as Add(-1)

 - Golden rules (or panic/deadlock)
   • Add() before starting the goroutine      (never after)
   • Exactly one Done() per Add(1)
   • Too many Done()      → panic
   • Add(+) after Wait()  → panic
   • Forget Done()        → Wait() waits forever (deadlock)

	- If you call Done() without calling Add() first, the WaitGroup will panic and crash the program.
	- If the WaitGroup counter is set to 5 but only three Done() calls happen, the counter never reaches zero, so any goroutine waiting on it will block indefinitely.

 - Easy mental picture
   There is one counter that says how many jobs are still running.
   Add(5) → counter = 5
   Every Done() → counter -= 1
   Wait() → keeps looking until counter == 0, then continues.

	 A sync.WaitGroup keeps everything in one 64-bit number. The top (high) 32 bits are the only thing we care about — they hold the count of unfinished tasks (the counter). When you call wg.Add(3), it puts 3 in those high 32 bits. Each wg.Done() subtracts 1 from that same high part. When the main goroutine runs wg.Wait(), it simply checks those high 32 bits. If the number there is still bigger than 0, it goes to sleep. As soon as the last Done() makes those high 32 bits exactly 0, Wait() wakes up instantly and the program continues to the next line.

	 High 32 bits = unfinished jobs.
	 High 32 bits == 0 → everything is done → keep going.

 - Deadlock analogy
   "I will pass you the ball only after you pass it to me.
    You will pass it only after I pass it to you."
   → Both wait forever → deadlock.

 - Inside the source – the 64-bit trick
   One single uint64 holds everything:

   bits 63–32 → counter        (unfinished tasks)   ← we care about this
   bits 31–12 → waiter count   (how many goroutines are inside Wait())
   bits 11–0  → tiny semaphore (used to wake up waiters)

   That's why you see in the source:
   counter := atomic.LoadUint64(&wg.state) >> 32   // throws away low 32 bits

   When counter finally becomes 0, the runtime releases the tiny semaphore
   and all sleeping Wait() calls wake up instantly.

 - Right shift >> (one-line memory)
   >> n = slide bits right n places → zeros fill from left → same as divide by 2ⁿ
   Example: value >> 32 → drops the low 32 bits → gives the real counter

 - Bonus facts
   • You can Add(100) once instead of 100 times
   • Never copy a WaitGroup (forbidden since Go 1.20+)
   • No way to read the current counter from outside (on purpose)
*/
