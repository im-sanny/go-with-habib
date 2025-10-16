# Class 4 – Context Switching, PCB, Concurrency

## Context Switching

- Context switching = CPU switching from one process to another.
- CPU can only execute one instruction at a time (single core).
- Because switching happens extremely fast (microseconds), it _appears_ to run multiple processes at once.
- OS saves the current process state and loads the next one — this is the “context switch”.

## Program Counter & Execution Flow

- Program Counter (PC) → holds address of next instruction.
- Control Unit (CU) → fetches the instruction from memory (RAM).
- Instruction Register (IR) → stores the fetched instruction.
- CU decodes the instruction, places operands in registers (AL, BL, CL, etc).
- ALU executes the operation.
- Result stored in data registers or memory.
- PC increments → points to next instruction.
- This fetch–decode–execute cycle continues until the program ends.

## CPU Speed

- A modern CPU can perform around 1–5 billion operations per second (1–5 GHz).

## Role of OS

- OS controls CPU scheduling and decides which process gets CPU time.
- OS updates which instruction address (PC value) will execute next.
- OS performs context switching between processes when needed (multitasking).

## PCB (Process Control Block)

- PCB = Process Control Block, a small data structure maintained by the OS.
- Contains all necessary info about a process:
  - Process ID (PID)
  - Process state (running, ready, waiting, etc.)
  - CPU registers (PC, SP, BP, etc.)
  - Memory info (code, data, stack)
  - I/O info and priority
- When a context switch occurs:
  → OS saves current process’s PCB (state)
  → Loads next process’s PCB (restores state)
- This allows processes to pause and resume smoothly.

## Process State

- “State” = the complete snapshot of CPU registers and memory for a process at a moment in time.

## PID

- PID = Process ID, a unique number used by OS to identify each process.
- Stored inside PCB.

## Concurrency

- Concurrency = multiple processes _appear_ to run simultaneously.
- In reality, CPU executes one process at a time but switches rapidly.
- Multicore CPUs can execute truly parallel tasks (parallelism).
- Concurrency → illusion of simultaneous execution through fast switching.

## Software

- Software = binary executable file (compiled instructions that CPU can execute).
- When executed, it becomes a process (loaded into memory).

## Summary

- Context Switching → CPU switches processes (OS saves/loads PCB).
- PCB → stores process info (PID, state, registers, memory, etc.).
- PC → next instruction address.
- Concurrency → illusion of multiple tasks running together.
- OS → manages CPU scheduling, context switching, and process states.
  _/
  /_
  ======================================
  Class 4 – Context Switching, PCB, Concurrency
  ======================================

## Context Switching

- Context switching = CPU switching from one process to another.
- CPU can only execute one instruction at a time (single core).
- Because switching happens extremely fast (microseconds), it _appears_ to run multiple processes at once.
- OS saves the current process state and loads the next one — this is the “context switch”.

## Program Counter & Execution Flow

- Program Counter (PC) → holds address of next instruction.
- Control Unit (CU) → fetches the instruction from memory (RAM).
- Instruction Register (IR) → stores the fetched instruction.
- CU decodes the instruction, places operands in registers (AL, BL, CL, etc).
- ALU executes the operation.
- Result stored in data registers or memory.
- PC increments → points to next instruction.
- This fetch–decode–execute cycle continues until the program ends.

## CPU Speed

- A modern CPU can perform around 1–5 billion operations per second (1–5 GHz).

## Role of OS

- OS controls CPU scheduling and decides which process gets CPU time.
- OS updates which instruction address (PC value) will execute next.
- OS performs context switching between processes when needed (multitasking).

## PCB (Process Control Block)

- PCB = Process Control Block, a small data structure maintained by the OS.
- Contains all necessary info about a process:
  - Process ID (PID)
  - Process state (running, ready, waiting, etc.)
  - CPU registers (PC, SP, BP, etc.)
  - Memory info (code, data, stack)
  - I/O info and priority
- When a context switch occurs:
  → OS saves current process’s PCB (state)
  → Loads next process’s PCB (restores state)
- This allows processes to pause and resume smoothly.

## Process State

- “State” = the complete snapshot of CPU registers and memory for a process at a moment in time.

## PID

- PID = Process ID, a unique number used by OS to identify each process.
- Stored inside PCB.

## Concurrency

- Concurrency = multiple processes _appear_ to run simultaneously.
- In reality, CPU executes one process at a time but switches rapidly.
- Multicore CPUs can execute truly parallel tasks (parallelism).
- Concurrency → illusion of simultaneous execution through fast switching.

## Software

- Software = binary executable file (compiled instructions that CPU can execute).
- When executed, it becomes a process (loaded into memory).

## Summary

- Context Switching → CPU switches processes (OS saves/loads PCB).
- PCB → stores process info (PID, state, registers, memory, etc.).
- PC → next instruction address.
- Concurrency → illusion of multiple tasks running together.
- OS → manages CPU scheduling, context switching, and process states.
