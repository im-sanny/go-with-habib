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

/*
Code segment:
-------------
• const p = 11
• printHello() { ... }
• main() { ... }

(Note: `var a` is not in the code segment, because it’s a variable
that can change during runtime.)

Data segment:
-------------
• var a = 10
• const p = 11
*/

/*
Now when we run the program, here’s what happens step by step:

1. The OS creates a **process** for this program.
   This process gets its own memory layout:
   - Code segment  → compiled Go functions
   - Data segment  → globals, constants
   - Heap          → dynamic memory + goroutine stacks
   - Stack         → for main thread (default 8MB in Linux)

2. The process starts with **one default OS thread** called the **main thread**.
   This thread is managed by the **kernel**, not by Go.
   The kernel executes threads using the CPU.

3. Inside this main thread, the **Go runtime** starts automatically.
   The runtime:
   - Initializes the **scheduler** (for running goroutines)
   - Initializes the **heap allocator**
   - Initializes the **garbage collector**
   - Creates **logical processors (P)**, by default equal to CPU cores

4. The Go runtime also creates:
   - **M (Machine)** → represents an OS thread (created by the kernel)
   - **P (Processor)** → logical context used by the scheduler
   - **G (Goroutine)** → lightweight function/task created by `go` keyword

   Together they form the **GMP model**:
   → Goroutines (G) run on Machines (M) using Processors (P).

5. The runtime first creates:
   - `G0` → a special internal goroutine for runtime work
   - `Gmain` → the main goroutine that runs your `main()` function

6. When `main()` runs and you write:
      go printHello(1)
      go printHello(2)
      ...
   Each `go` statement creates a **new goroutine**.
   These goroutines are very lightweight:
   - Each starts with ~2KB stack in the **heap** (not on the main stack)
   - The stack can grow and shrink automatically
   - The scheduler decides when and where each goroutine will run

7. The **kernel doesn’t know about goroutines**.
   It only knows the OS threads (M) created by Go runtime.
   The Go scheduler runs in user space and decides which goroutine (G)
   runs on which OS thread (M).

   So the real chain is:
   → CPU executes thread instructions
   → Kernel schedules OS threads
   → Go runtime scheduler runs goroutines inside those threads

8. The line `time.Sleep(5 * time.Second)` keeps the main goroutine alive
   long enough so that all other goroutines can finish their execution.

9. Finally, when all goroutines are done, the program exits,
   and the OS cleans up the process and memory.
*/

/*
Extra summary:
--------------
• Main thread stack size: 8 MB (by OS)
• Goroutine stack size: starts at ~2 KB (in heap)
• Code & data segments are loaded once per process
• Go runtime, heap, GC, scheduler all live inside the process
• Kernel controls only OS threads, not goroutines
• CPU executes instructions from threads assigned by the kernel
*/
