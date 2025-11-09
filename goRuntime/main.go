package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "hello go runtime!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I'm Sanny! Junior Software Engineer.")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting on the server", err)
	}
}

/*
- When the OS boots, RAM is divided into two main parts — kernel space (for the operating system) and user space (for applications).
- Process and process-related things run from and stay in user space.
- A system call (syscall) is the way a user-space process requests a service from the kernel. Since user-space programs can’t directly access hardware or kernel functions (for safety), they have to ask the kernel to do it for them — and that’s done through a system call.
- If a user-space process needs access to any file, socket, or network it requests the kernel to provide it.
- Software running in user space can’t access or modify anything in kernel space if the kernel doesn’t allow it.
- A process runs isolated; it doesn’t know anything about what’s around it.
- Kernel manages and coordinates everything happening in user space.


Go Runtime:
- Go runtime: The Go runtime is the core engine of Go that manages how Go programs run. It acts like a mini operating system inside your program — handling goroutine scheduling, memory management, garbage collection, and system interactions such as network I/O and threading.

One liner: The Go runtime is the core engine of Go that manages goroutines, memory, scheduling, and system-level tasks within a Go program.

- Go runtime runs first when we execute Go code — even before any init() function.

- The OS executes the startup assembly code (part of Go’s runtime). That code initializes the Go runtime. Then the runtime executes all init() functions and finally your main().

- Go code  execution chain: CPU → OS → Process → Main Thread → Go Runtime → Go Code

Go Runtime Does:
1. Initialize Go Scheduler – sets up goroutine scheduling, logical processors (P), and OS threads (M).
2. Syscalls & OS Integration – makes syscalls to the kernel, e.g., request epoll_create for network polling.
3. Epoll Management – handles epoll_create, epoll_ctl, and epoll_wait.
4. Garbage Collection (GC) Setup – initializes and runs the garbage collector.
5. Stack Management – manages goroutine stack growth and shrinking.
6. Timers & Concurrency Primitives – sets up timers and basic primitives needed for goroutines.

go runtime -> syscall -> kernel -> epoll_create
epoll: A Linux kernel mechanism that efficiently monitors multiple file descriptors and notifies when I/O events (like read/write) are ready.


Imagine the kernel creates 3 processes, and each has 2 threads — total 6 threads.
The CPU (not the kernel) runs them, switching quickly between threads.
Suppose one thread wants to read from a file.
1. The thread makes a system call to the kernel.
2. The kernel checks its file descriptor table and gives back a file descriptor (FD) — a small ID representing the file.
3. If the file’s data isn’t ready, the thread uses epoll:
- It calls epoll_ctl() to tell the kernel “I’m waiting for this FD.”
- Then calls epoll_wait() — and the kernel makes the thread sleep.
4. While that thread sleeps, the CPU runs other threads.
5. When the kernel sees the file is ready, it wakes the thread, and epoll_wait() returns.
6. The thread then calls read() to get the actual data.

- in linux it call epoll, in mac it's kqueue, in ewindows it's iocp
- epoll is a Linux system call that helps programs efficiently monitor multiple file descriptors (like network sockets, files, pipes) to see if they're ready for I/O operations (reading or writing).

Go Runtime:
1. Go runtime starts → It initializes the scheduler (M, P, G model).
2. It then creates an epoll instance (via `epoll_create1` syscall).
3. It also starts a separate OS thread called the Netpoller Thread.
   * This thread runs epoll_wait() in a loop to watch for network I/O events.
4. Now when any goroutine tries to read/write on a socket that isn’t ready:
   * The Go runtime calls epoll_ctl() (syscall to kernel) to register that socket’s FD with epoll.
   * Then it parks the goroutine (makes it sleep).
   * The OS thread that was running that goroutine becomes free for other goroutines.
5. When the kernel detects data ready on that FD, it wakes the Netpoller Thread via epoll_wait().
6. The Netpoller Thread then notifies the Go runtime that “FD 7 is ready.”
7. The runtime then marks the waiting goroutine runnable again, schedules it to run, and when it resumes, the goroutine does the actual read() syscall to fetch the data.

Netpoller Thread: The Netpoller Thread in Go is basically a special OS thread that the Go runtime creates to efficiently handle network I/O using epoll (on Linux) or similar mechanisms on other OSes.

GC setup:
Go runtime initializes the garbage collector during startup.
It doesn’t create a special thread for GC — instead, GC runs on the existing OS threads (M’s) managed by the Go scheduler.
GC works concurrently with other goroutines to clean memory in the background.

Go runtime uses the G-P-M model to schedule goroutines on OS threads.
Goroutine (G) → Processor (P) → OS Thread (M) → CPU

G = your Go code/task/ goroutine
P = scheduler context/run queue/logical processors
M = real OS thread/machine thread
CPU = executes M

Each P (logical processor) has one local run queue, and that queue can hold up to 256 goroutines.

There are two types of run queues:

Local run queue – one per P (fixed size, 256 slots).
Global run queue – shared by all Ps, dynamic in size (limited only by available memory).

Go Scheduler:
Suppose I have 4 logical processors (P1–P4).
Each logical processor has one local run queue, and each queue has 256 slots.
Each slot can hold one goroutine, so up to 256 goroutines can be queued per P (if that many exist).

Each P is attached to an OS thread (M).
The OS threads are what the CPU cores actually execute.

When a logical processor runs, it picks goroutines from its local run queue one by one and executes them.
If multiple CPU cores are available, then multiple Ps (and thus multiple goroutines) can run in parallel.
If there’s only one CPU core, they run concurrently (time-shared).

Now about the global run queue:
If all 4 logical processors (P1–P4) have their local run queues full (256 slots each) and a new goroutine is created, there’s no space left locally to place it.
In that case, the extra goroutine is placed into the global run queue.

Later, if one of the processors finds its local run queue empty, it will first try to steal goroutines from the local queues of other processors (this is called work stealing).
If it can’t steal any, it will then take goroutines from the global run queue (if available) and put them into its local queue to execute.

*/
