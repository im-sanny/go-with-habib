## Threads:

- Threads are units of execution inside a process. They help execute the program that a process gets from the OS.

## Analogy:

- Process is a kitchen in a restaurant, and a thread is a chef who works in the kitchen and cooks.
- Two chefs in the same kitchen can easily use everything in their kitchen. But if they want to use items from another kitchen, it’s not easy. Similarly, multiple threads of the same process can share resources easily, but each process is separate, so it cannot directly access another process’s data.

## Thread Creation:

- When a process is created, it has 1 thread by default.
- A process can create multiple threads.
- Threads help execute one or many programs from a process.
- A process creates a thread to execute the program assigned by the OS. So, the thread does the execution, but it belongs to the process.

## Process vs Thread:

- Process: has its own memory (code, data, heap, stack) and at least one thread.
- Thread: shares process memory, but has its own stack, registers, and program counter.

## Execution Flow:

- Process comes first → threads are created inside it → CPU executes threads.
- Logical CPU, virtual CPU, or core is a hardware thread, and a software thread is an execution unit inside a process.
- A process with multiple threads can have those threads run simultaneously on different vCPUs. This allows parallel execution and faster performance when vCPUs are available.

## OS Role:

- OS is the bridge between hardware and software.
- The OS orchestrates but doesn't directly manipulate the PC in normal operation.
