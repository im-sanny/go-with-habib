## What epoll does:

Suppose the kernel created 3 processes, and each process has 2 threads — so in total there are 6 threads. The kernel itself doesn’t run them; the CPU does. If it’s a 1-core CPU with 2 logical processors, these 2 logical CPUs can run the 6 threads concurrently (switching between them very fast).
Now, one of these threads wants to read a file. It makes a system call to the kernel. The kernel checks the file descriptor table and returns a file descriptor (FD) — a token/ID that represents the file. This FD verifies the thread’s request and tells the kernel exactly which file to access.
If the file or data isn’t immediately ready, the kernel can’t give it to the thread right away. So the thread can use epoll: it calls epoll_ctl to register interest in that FD, then calls epoll_wait(). This puts the thread to sleep until the kernel marks the FD as ready.
While the thread sleeps, the CPU runs other threads. When the kernel finally gets the file data or the event is ready, it wakes up the thread via epoll_wait() and returns the FD. The thread then uses that FD to send a read request to the kernel, which allows the thread to read the actual file data.

## Go runtime and Netpoller thread:

- Go runtime: At startup, Go runtime creates an epoll instance and spawns a "Netpoller Thread" that runs epoll_wait to monitor file descriptors. When a goroutine tries to read from a socket but data isn't ready, Go parks the goroutine and registers the file descriptor with epoll using epoll_ctl, freeing the OS thread for other work. The Netpoller Thread sleeps in epoll_wait until the kernel signals that data is ready on one or more sockets. Once notified, the runtime identifies which parked goroutines were waiting for those file descriptors, marks them as runnable, and schedules them for execution. The goroutines resume and complete their I/O operations. This design allows Go to efficiently handle thousands of concurrent connections without blocking threads, since parked goroutines don't consume OS thread resources while waiting.

- Netpoller Thread: The netpoller thread is an internal Go runtime mechanism that uses OS-level async I/O APIs to monitor network sockets for readiness, enabling efficient, scalable, non-blocking I/O while allowing Go code to use simple blocking-style syntax

- Short version:
  At startup, Go runtime creates an epoll instance and starts a Netpoller Thread to watch sockets.
  When a goroutine tries to read from a socket but the data isn’t ready: 1. Go parks the goroutine (puts it aside). 2. Registers the socket’s file descriptor with epoll (epoll_ctl) so the kernel will notify when data is ready. 3. The OS thread running that goroutine is now free to do other work.
  The Netpoller Thread waits in epoll_wait() until the kernel signals that data is ready.
  When data is ready, the runtime: 1. Finds which goroutines were waiting for those sockets. 2. Marks them runnable. 3. Schedules them to continue their I/O.
  This way, thousands of goroutines can wait for I/O without blocking OS threads, making Go efficient for concurrent connections.
