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


Suppose the kernel created 3 processes, and each process has 2 threads — so in total there are 6 threads. The kernel itself doesn’t run them; the CPU does. If it’s a 1-core CPU with 2 logical processors, these 2 logical CPUs can run the 6 threads concurrently (switching between them very fast).
Now, one of these threads wants to read a file. It makes a system call to the kernel. The kernel checks the file descriptor table and returns a file descriptor (FD) — a token/ID that represents the file. This FD verifies the thread’s request and tells the kernel exactly which file to access.

If the file or data isn’t immediately ready, the kernel can’t give it to the thread right away. So the thread can use epoll: it calls epoll_ctl to register interest in that FD, then calls epoll_wait(). This puts the thread to sleep until the kernel marks the FD as ready.
While the thread sleeps, the CPU runs other threads. When the kernel finally gets the file data or the event is ready, it wakes up the thread via epoll_wait() and returns the FD. The thread then uses that FD to send a read request to the kernel, which allows the thread to read the actual file data.

- in linux it call epoll, in mac it's kqueue, in ewindows it's iocp
- epoll is a Linux system call that helps programs efficiently monitor multiple file descriptors (like network sockets, files, pipes) to see if they're ready for I/O operations (reading or writing).

- epoll: When a thread requests the kernel to read a file, the thread calls epoll_ctl to register that file descriptor with epoll. The kernel then makes the thread sleep so it doesn't waste CPU time. While the thread sleeps, the kernel prepares the file data. When the data is ready, the kernel marks that file descriptor as ready. The sleeping thread then calls epoll_wait and the kernel returns which file descriptor is ready (let's say FD 7). The thread wakes up and makes a read() request to the kernel using that file descriptor. The file descriptor works as a token/ID to tell the kernel which file to access, and the kernel lets the thread read the data from that file.

- Go runtime: At startup, Go runtime creates an epoll instance and spawns a "Netpoller Thread" that runs epoll_wait to monitor file descriptors. When a goroutine tries to read from a socket but data isn't ready, Go parks the goroutine and registers the file descriptor with epoll using epoll_ctl, freeing the OS thread for other work. The Netpoller Thread sleeps in epoll_wait until the kernel signals that data is ready on one or more sockets. Once notified, the runtime identifies which parked goroutines were waiting for those file descriptors, marks them as runnable, and schedules them for execution. The goroutines resume and complete their I/O operations. This design allows Go to efficiently handle thousands of concurrent connections without blocking threads, since parked goroutines don't consume OS thread resources while waiting.

- Netpoller Thread: The netpoller thread is an internal Go runtime mechanism that uses OS-level async I/O APIs to monitor network sockets for readiness, enabling efficient, scalable, non-blocking I/O while allowing Go code to use simple blocking-style syntax

GC setup:
Go runtime initializes the garbage collector during startup.
It doesn’t create a special thread for GC — instead, GC runs on the existing OS threads (M’s) managed by the Go scheduler.
GC works concurrently with other goroutines to clean memory in the background.

Go runtime uses the G-P-M model to schedule goroutines on OS threads.
Goroutine (G) → Processor (P) → OS Thread (M) → CPU

G = your Go code/task
P = scheduler context/run queue
M = real OS thread
CPU = executes M

g = goroutine
p = logical processors
m = os thread/machine thread

Each P (logical processor) has one local run queue, and that queue can hold up to 256 goroutines.

There are two types of run queues:

Local run queue – one per P (fixed size, 256 slots).
Global run queue – shared by all Ps, dynamic in size (limited only by available memory).


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
