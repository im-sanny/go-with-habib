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

	fmt.Println("Computer architecture!")
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

	each cell address  usually starts from 0, 1, 2.
	if a computer is 8 bit then it's cell keeps 8 byte in a cell of ram.

	if a computer is 16 bit then its cell address starts like 0, 2, 4, 6, 8, 10, 12, 14, 16 -> each cell keeps 2 byte
	if a computer is 32 bit then its cell address starts like 0, 4, 8, 12 .... -> each cell keeps 4 byte
	if a computer is 64 bit then its cell address starts like 0, 8, 16, 24 .... -> each cell keeps 8 byte

Always the value of the SP will be lower than BP.

stack starts to execute from bottom or lower of it's address of space
*/
