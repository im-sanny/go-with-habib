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
	Go Runtime:
- Go runtime: The Go runtime is the core engine that powers how Go programs run.
It’s like a mini operating system inside your Go program, managing everything that happens when your code executes — especially concurrency, memory, and scheduling.

- When the OS boots, RAM is divided into two main parts — kernel space (for the operating system) and user space (for applications).
- Process and process-related things run from and stay in user space.
- A system call (syscall) is the way a user-space process requests a service from the kernel. Since user-space programs can’t directly access hardware or kernel functions (for safety), they have to ask the kernel to do it for them — and that’s done through a system call.
- If a user-space process needs access to any file, socket, or network it requests the kernel to provide it.
- Software running in user space can’t access or modify anything in kernel space if the kernel doesn’t allow it.
- A process runs isolated; it doesn’t know anything about what’s around it.
- Kernel manages and coordinates everything happening in user space.

- The OS executes the startup assembly code (part of Go’s runtime). That code initializes the Go runtime. Then the runtime executes all init() functions and finally your main().

go runtime will do:
1. initialize go schedular
2. go runtime will syscall to kernel then it'll request the kernel for creating a epoll_create for it.
go runtime -> syscall -> kernel -> epoll_create
3. epoll: epoll is a feature of the kernel(os), epoll does 3 operation: epoll_create, epoll_ctl, epoll_wait
4. setup GC

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


*/
