package main

import "fmt"

func add(x int, y int) int {
	res := x + y
	return res
}

func main() {
	var a int = 10

	sum := add(a, 4)
	fmt.Println(sum)

}

/*
Computer Architecture:

Motherboard → Connects CPU, RAM, storage, I/O.
CPU → Brain of computer.
 ALU = calculations + logic.
 CU & Registers = control + fast storage.
RAM → Volatile, temporary memory (8-bit = 1 byte per cell).
Storage (HDD/SSD) → Permanent, non-volatile.
Bus System → Data, Address, Control buses.
Generations of Computers:
 1.Vacuum tubes (1940s)
 2.Transistors (1950s)
 3.Integrated Circuits (1960s)
 4.Microprocessors (1970s–Now)
 5.AI / Quantum (Future).
if a computer is 8 bit then each cell of the RAM can keep 8 bit of data, so each cell can hold 0 to 256 in decimal number.

Computer History:

Abacus → First known tool.
1703 Leibniz → Binary system.

1830s Babbage → Analytical Engine (first concept of general-purpose computer).
Ada Lovelace → First programmer.
Alan Turing → Father of CS & AI, Turing Machine.
1945 ENIAC → First electronic computer, vacuum tubes.
1947 Transistor → Replaced vacuum tubes (Bell Labs).
1958 IC (Kilby, Noyce) → Multiple circuits on a chip.
1971 Intel 4004 → First microprocessor (4-bit).
1981 IBM PC → Popularized personal computing (16-bit).
*/

/*
OS and Computer Architecture:

1. OS (Operating System) → Bridge between hardware & software.
2. Main role → Runs programs, manages CPU, memory, files, and I/O.
3. When system starts →
	  i. BIOS/UEFI runs → loads OS from HDD/SSD to RAM.
	  ii. CPU executes OS + user programs.r7r[7]y'y
4. RAM → Temporary, fast, volatile memory.
5. HDD/SSD → Permanent storage.
6. CPU → Works continuously until system off.

Early Computing & ENIAC:

ENIAC (1945) → First general-purpose electronic computer.
Used punch cards for input, output, and storage.
IBM 80-column punch cards, 12 rows each.
No OS — everything done manually by operators.
OS idea came from automating operator’s manual work.

1956 ->	GM-NAA I/O – first OS (batch system)
1960s ->	Keyboard input, time-sharing OS
1970s ->	UNIX developed
1980s ->	macOS, MS-DOS (personal computers)
1990s ->	Linux (open-source UNIX-like)
2000s -> Now GUI-based OS (Windows, macOS, Linux, Android, iOS)
*/

/*
CPU & Processing:

• CPU = brain of computer → performs all logic & control.
• CPU parts → Processing Unit (PU) + Register Set.
• PU = ALU (calculations, logic) + CU (control, instruction flow).
• Registers = tiny, fast memory inside CPU.
• SP → top of stack | BP → base of stack frame.
• IR → current instruction | PC → next instruction address.
• AX, BX, CX, DX → general purpose registers.
• 8-bit CPU → AL, BL, CL, DL (1 byte each).
• 16-bit CPU → AX = AH (high 8 bits) + AL (low 8 bits).
• 32-bit → EAX etc | 64-bit → RAX etc.
• 8-bit CPU ⇒ each RAM cell = 8 bits (1 byte).
• Process = program in execution (using RAM + CPU time).
• Each process has its own memory (code, data, stack, heap).
• Stack stores temp data for functions & local vars.
• SP & BP manage stack frames.
• Summary → CPU = PU + Registers | PU = ALU + CU | SP/BP = stack control.

A process behaves like a virtual computer because it has its own virtual CPU, memory space (RAM), and other resources provided by the operating system.
*/

/*
SP vs BP:

Memory Architecture:
- 8-bit: addresses 0, 1, 2... (byte-addressable)
- 16-bit: addresses 0, 2, 4... (but still byte-addressable)
- 32-bit: addresses 0, 4, 8...
- 64-bit: addresses 0, 8, 16...

STACK BEHAVIOR:
- Stack grows DOWNWARD (from high to low addresses)
- SP always points to the CURRENT TOP of stack
- BP points to a STABLE BASE within current function frame
- BP is typically at HIGHER address than SP
- SP is dynamic; BP is relatively stable during function execution

STACK FRAME LAYOUT (growing downward):
[HIGHER ADDRESSES]
... Caller's data ...
[Parameter N]
...
[Parameter 1]
[Return Address]
[Saved BP]        ← BP points here (function frame start)
[Local Variable 1]
...
[Local Variable N] ← SP points here (current stack top)
[LOWER ADDRESSES]

FUNCTION CALL PROCESS:
1. Caller pushes arguments (right to left)
2. Caller pushes return address
3. CALL instruction jumps to function
4. Function prologue:
   - Push current BP (save caller's frame)
   - Move SP to BP (set new frame base)
   - Adjust SP downward for local variables
5. Function executes
6. Function epilogue:
   - Move SP to BP (clean locals)
   - Pop old BP (restore caller's frame)
   - RET instruction pops return address and jumps back

MEMORY SEGMENTS:
- Code Segment: Program instructions
- Data Segment: Global/static variables
- Stack: Function calls, local variables (grows downward)
- Heap: Dynamic allocation (grows upward)

Return Address:
1. RETURN ADDRESS FIRST:
   - When CALL instruction executes, the CPU automatically:
     1. Pushes return address (where to come back to)
     2. Jumps to the new function
   - This ensures the function knows how to return

2. SP MOVEMENT:
   - SP always moves in fixed increments based on system:
     - 32-bit: ±4 bytes each push/pop
     - 64-bit: ±8 bytes each push/pop
   - SP moves DOWN when pushing (adding to stack)
   - SP moves UP when popping (removing from stack)
   - Example (32-bit):
     Start: SP = 1000
     Push: SP = 996 (1000 - 4)
     Push: SP = 992 (996 - 4)
     Pop:  SP = 996 (992 + 4)

3. TYPICAL FUNCTION CALL SEQUENCE:
   - CALL pushes return address
   - Function pushes old BP (save caller's frame)
   - Function sets new BP = SP (create new frame)
   - Function adjusts SP downward for local variables

KEY POINTS:
- Each function call creates a new stack frame
- BP provides stable reference to access parameters and locals
- SP tracks the current stack top for push/pop operations
- Stack frames are nested (like Russian dolls)
- When function returns, its stack frame is effectively "discarded"
*/

/*
Context Switching, PCB, Concurrency:
Context switching: cpu switches from one process to another so fast that it looks like all running at once.

Program Counter (PC): points current instruction and moves to next after each execution.
Control Unit (CU): takes instruction from memory, decodes it, and sends to ALU for execution.
ALU: does calculation or logic part and stores result in register.
Modern CPU: can do around 1 billion operations per second.
OS: controls cpu and decides which process runs next by managing PC and process states.

PCB (Process Control Block): small data box where OS keeps process info like id, state, registers, etc.
State: all the cpu info needed to continue process from where it stopped.

PID: unique process id saved inside PCB.
Concurrency: cpu runs one thing at a time but so fast that it looks like many are running together.
concurrency is a feelings.
Software: binary executable file that cpu can run as process.
*/

/*
Context Switching, Concurrency & Parallelism (Quick Note):

context switching → one hand scratching 2 places fast, one by one.
concurrency → one hand switching fast between tasks, all seem to progress.
parallelism → two hands scratching both places at same time.

context switching = cpu switching between processes.
concurrency = many tasks handled together (by fast switching). concurrency is illusion of parallelism, it's actually context switching but the process is so fast that we feel like parallelism.
parallelism = many tasks run truly together (multi-core).

core → physical cpu unit or small box inside the main cpu.
logical cpu (hardware thread) → virtual cpu inside core.
process = running program.
thread = smaller part inside process.

intel calls logical cpu as “threads” because each runs one software thread.

old cpu → only concurrency (context switching).
modern cpu → concurrency + real parallelism.

context switch = extra time (save/load process).
too many switches = slower performance.

scheduling algorithm → helps cpu choose what to run next.

types:
- FCFS → first come, first serve
- SJF → shortest job first
- RR → equal time slice
- Priority → high priority first
- Multilevel Queue → mixed methods
*/

/*
Example: After turning on the system, the CPU loads the OS data from the HDD into memory, starts executing it, and the OS comes into power. When we click on a music player, the OS creates a process to execute the program and play music. In the music process, there will be a code segment, data segment, stack, and heap. The process creates multiple threads for different tasks (UI thread, audio thread, file thread). Suppose the code segment has 10 lines of code. The program counter points to the first line, loads it into the instruction register, then increases the program counter by 1 to move to the next line or instruction. The control unit decodes the instruction, the arithmetic logic unit executes it, and it uses the register set if needed. Then it goes back to the program counter to run the next instruction. The CPU switches rapidly between threads using context switching.
*/
