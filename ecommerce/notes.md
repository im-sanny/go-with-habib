    ┌──────────────────────────────────────────────┐
    │           CLIENT                             │
    │  Browser → Sends HTTP request (GET /hello)   │
    └──────────────┬───────────────────────────────┘
                   │
                   ▼
    ┌──────────────────────────────────────────────┐
    │           NETWORK (Router etc.)              │
    │  Forwards packet to server’s IP:3000         │
    └──────────────┬───────────────────────────────┘
                   │
                   ▼
    ┌──────────────────────────────────────────────┐
    │           HARDWARE LAYER (Server NIC)        │
    │  - Converts EM signals → binary              │
    │  - Stores data in NIC receive buffer         │
    │  - Interrupts the kernel                     │
    └──────────────┬───────────────────────────────┘
                   │
                   ▼
    ┌──────────────────────────────────────────────┐
    │           OS / KERNEL NETWORK STACK          │
    │  - Moves data → socket receive buffer        │
    │  - Marks socket FD as readable               │
    │  - Uses epoll/kqueue to notify Go runtime    │
    └──────────────┬───────────────────────────────┘
                   │
                   ▼
    ┌──────────────────────────────────────────────┐
    │           GO RUNTIME LAYER                   │
    │  - Main goroutine blocked in Accept()        │
    │  - Wakes up when data available              │
    │  - Accepts new conn → spawns handler goroutine│
    │  - Scheduler manages all goroutines          │
    └──────────────┬───────────────────────────────┘
                   │
                   ▼
    ┌──────────────────────────────────────────────┐
    │           APPLICATION LAYER (Your Code)      │
    │  - ServeMux routes request (/hello, /about)  │
    │  - Handler executes and writes response      │
    │  - Data written to socket send buffer        │
    └──────────────┬───────────────────────────────┘
                   │
                   ▼
    ┌──────────────────────────────────────────────┐
    │           KERNEL + NIC (Sending)             │
    │  - Moves data to NIC transmit buffer         │
    │  - NIC sends EM signal → router → client     │
    └──────────────┬───────────────────────────────┘
                   │
                   ▼
    ┌──────────────────────────────────────────────┐
    │           CLIENT                             │
    │  - Receives response                         │
    │  - Browser renders "Hello world!"            │
    └──────────────────────────────────────────────┘
