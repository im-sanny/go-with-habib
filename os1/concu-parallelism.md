Context Switching, Concurrency, and Parallelism:

context switching:
like a man with one hand scratching two places — not at once, but switching fast between them.
cpu does the same — switches between processes very fast, one after another, so it looks like multitasking.

concurrency:
like a man with one hand doing multiple tasks by switching fast,
so even though only one thing happens at a time, everything seems to progress together.
it’s about managing many tasks at once (not necessarily running them at the same time).

parallelism:
like a man with two hands scratching two places at the same time.
cpu with multiple cores can truly run many things together at once without switching.

Concept:
Suppose Intel Core i3 CPU has 3 cores. Each core has 2 logical or virtual CPUs (called threads).
These logical CPUs aren’t real hardware but act like separate CPUs to handle tasks faster.
So, 3 cores × 2 threads = 6 logical CPUs total.

core = physical processing unit inside cpu or Each core is like a small home or box inside the main CPU.
logical cpu = virtual cpu inside core (created by hyper-threading).

process ≠ cpu.
process = a program running in memory.
cpu = the machine that runs those processes.

Modern CPUs have both concurrency and parallelism.
Old CPUs only had concurrency (by context switching), no real parallelism.

if cpu has 6 logical cpu, it can run 6 threads (or processes) truly at once (parallelism).
if there are 7 processes, then 6 will run at once and the extra one will wait or run by context switching (which creates concurrency).

process vs thread:
a process can have multiple threads inside it.
threads share the same memory of the process but run different parts of the task.
intel calls logical cpu as "threads" because each logical cpu can run one software thread at a time.
so they use the same word for hardware threads.

older cpus had only one core, so they only did context switching and concurrency.
modern cpus can do both concurrency and real parallelism.

context switching takes time because cpu has to save the current process state and load the next one.
so if there are too many switches, total execution time increases.

sometimes, running one after another can be faster because switching time is saved,
but cpu uses smart scheduling to balance speed and fairness.

scheduling algorithm:
it helps cpu decide which process or thread to run next and for how long.

some common scheduling algorithms:

1. FCFS (First Come First Serve) – runs in order of arrival.
2. SJF (Shortest Job First) – shorter tasks first.
3. RR (Round Robin) – gives each process equal small time slices.
4. Priority Scheduling – runs higher priority tasks first.
5. Multilevel Queue – mixes multiple algorithms for different job types.
