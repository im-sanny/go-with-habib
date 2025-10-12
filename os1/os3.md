### Class 3 – CPU and Processing

## CPU (Central Processing Unit)

- Brain of the computer, performs all calculations & control.
- Two main parts:
  1. Processing Unit (PU)
  2. Register Set

## Processing Unit (PU)

- ALU (Arithmetic Logic Unit): does +, -, ×, ÷, AND, OR, NOT, XOR.
- CU (Control Unit): fetches, decodes, and executes instructions.

## Register Set

- Very small & fast memory inside CPU for temporary data.
- Common Registers:
  SP (Stack Pointer): points top of the stack.
  BP (Base Pointer): points base of current stack frame.
  IR (Instruction Register): holds current instruction.
  PC (Program Counter): address of next instruction.
  AX, BX, CX, DX: general purpose registers (data transfer, calc etc).

## General Purpose Registers by CPU Bit

    8-bit:  AL, BL, CL, DL
    16-bit: AX(AH+AL), BX, CX, DX
    32-bit: EAX, EBX, ECX, EDX
    64-bit: RAX, RBX, RCX, RDX

- 8-bit CPU → each RAM cell = 8 bits (1 byte).
- 16-bit CPU → AX split into AH (high) + AL (low).

## Process

- A program in execution (from start to end).
- Uses part of RAM and CPU time.
- Each process has its own memory space (code, data, heap, stack).

## Stack Management

- Stack used for function calls, local vars, return addresses.
- SP & BP manage stack frame:
  SP → top of stack
  BP → base of current function’s frame

## Summary

CPU = PU + Register Set
PU = ALU + CU
Registers = SP, BP, IR, PC, AX–DX
Process = running program
SP & BP = manage stack frames
