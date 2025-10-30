package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func printHello(num int) {
	fmt.Println("hello fumis", num)
}

func main() {
	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)

	fmt.Println(a, "", p)

	time.Sleep(5 * time.Second)
}

/*
Goroutine:
- Go's concurrent, lightweight execution unit.
- Go Runtime works like a virtual operating system.
- When a Go program runs, it runs its own mini OS called the Go Runtime.
- Goroutine is a lightweight thread (also called virtual or logical thread).
- It can execute many functions concurrently.
- Go Runtime manages creation, scheduling, and execution of goroutines.
- Before any function, if you put the keyword 'go', it becomes a goroutine.
- Go Runtime = Mini OS for running Go code.

Main Goroutine:
- If we use a goroutine inside main(), then that main function becomes the Main Goroutine.
- The main thread starts and runs the Go Runtime.
- The Go Runtime then executes the main goroutine.
- When the main goroutine finishes, the whole program finishes.
- To keep other goroutines running, we must keep the main goroutine alive (for example: using WaitGroup, channel, or time.Sleep).

Thread vs Goroutine:
- Programmer creates goroutines, not OS threads.
- OS threads are created by the Go Runtime depending on how many goroutines need to run.
- OS thread stack size ≈ 8MB.
- Goroutine stack size starts with ≈ 2KB and grows dynamically when needed.
- Go Runtime copies old stack data to a new larger stack when it grows and frees the old one.
- Normal OS threads are managed by OS kernel, goroutines are managed by Go Runtime.
- Each goroutine’s stack stays in heap memory.

Go Runtime Initialization:
When Go Runtime starts, it initializes the following:
    1. Goroutine Scheduler
    2. Heap Allocator
    3. Garbage Collector
    4. Logical Processors

Main Thread and Go Runtime:
- Main thread executes Go Runtime.
- Go Runtime initializes its subsystems and runs the main goroutine.

Go Routine Scheduler:
- OS Kernel scheduler handles process and OS thread scheduling.
- Go Scheduler handles goroutine scheduling inside the Go Runtime.
- It works similarly to a kernel scheduler but runs in user space.
- The Go Scheduler decides which goroutine runs on which OS thread and when.
- It can run thousands of goroutines efficiently by mapping them to a few OS threads.
- Scheduler continuously balances workloads across available threads.

Logical Processors (P):
- Inside OS, multiple vCPUs (virtual CPUs) exist for each physical CPU core.
- Go Runtime creates the same number of Logical Processors (P) as vCPUs by default (GOMAXPROCS).
- Each Logical Processor (P) is attached to an OS thread (M).
- Each OS thread executes goroutines (G) using its assigned processor (P).
- Example:
    CPU: 2 cores → 4 vCPUs
    Go Runtime creates 4 Logical Processors.
    OS creates 4 OS Threads for these processors.
    Total threads in process = 4 OS threads + 1 Main thread = 5 threads.
    OS Kernel tracks 5 threads.
- Go Scheduler schedules goroutines to run on these OS threads through logical processors.
- 4 threads can execute 10 or more goroutines concurrently (managed by scheduler).

G-M-P Model:
    G → Goroutine (actual code/task)
    M → Machine (OS thread)
    P → Processor (logical processor that schedules goroutines)
- Each P has its own run queue of goroutines.
- Scheduler picks goroutines from P’s run queue and assigns them to available M.
- Go Runtime maps many Gs → few Ms using Ps efficiently.

Go Runtime as Virtual OS:
- Go Runtime works as a kernel for goroutines.
- This kernel has its own scheduler.
- Scheduler divides goroutines’ execution work among OS threads.
- OS threads execute goroutines from CPU.
- Go Scheduler decides which goroutine runs and when.
- Go Runtime can be considered a Mini OS or Virtual Operating System for goroutines.

Execution Flow Summary:
    Thread execute → Process start
    Process → Virtual Computer
    Go Runtime → Virtual Operating System
    Process start → Go Runtime execute
    Go Runtime → Scheduler + Heap Allocator + Garbage Collector + Logical Processors
    Scheduler → Executes goroutines using OS threads.

Effects of Excessive Goroutines:
- Scheduler notices when too many goroutines are created.
- Go Runtime may create more OS threads or logical processors if needed (up to limits).
- If system memory (RAM) is full, Go Runtime cannot create new OS threads.
- ❌ No new OS threads → ❌ Some goroutines cannot execute.
- Too many goroutines increase memory use and scheduling overhead.

Goroutine's Home: Stack & Heap:
- Every goroutine’s stack lives in the heap.
- Each goroutine starts with 2KB stack.
- If 2KB isn’t enough, Go Runtime dynamically increases stack size.
- Go Runtime copies the current stack to a new bigger one and updates stack pointers (SP, BP, return addresses).
- Stack frame creation and execution happen per goroutine in its own stack.
- main() → Main Goroutine.
- Writing 'go' before any function makes it a new goroutine.

In short:
    - Go Runtime creates and manages goroutines.
    - Goroutines are lightweight, concurrent execution units.
    - Scheduler runs many goroutines efficiently on few OS threads.
    - Go Runtime acts as a mini OS inside the process.
    - Main thread runs the Go Runtime → Go Runtime runs the Main Goroutine.
*/
