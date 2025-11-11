package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Sanny, Junior software engineer.")
}

func main() {
	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler) //route

	mux.HandleFunc("/about", aboutHandler) //route

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting on the server", err)
	}
}

/*
-	A file descriptor is an integer that represents an open file.
- File descriptors are integers that start from 0 — where 0, 1, and 2 represent stdin, stdout, and stderr.
- A socket is a communication endpoint used to send and receive data over a network.
- In Unix and Linux, everything is treated as a file — meaning all input/output operations are handled in the same way.
- When a server starts, one of the first things it does is create a socket to listen for incoming connections.

Short version of the process:
 Client request → NIC (waves to binary) → Kernel (routes to port 3000) → Go runtime (creates goroutine) → Handler executes (writes response) → Kernel → NIC (binary to waves) → Client receives response.
*/

/*
	When a client sends an HTTP request to the server, the request first travels through the network and reaches the server’s Network Interface Card (NIC). The NIC converts the incoming electromagnetic signals into binary data and stores it temporarily in the RAM, specifically inside the NIC receive buffer.

	After this, the NIC triggers an interrupt to notify the kernel that new data has arrived. The kernel then reads the data from the NIC buffer and moves it to another area in RAM called the socket receive buffer, which is associated with the port the server is listening on — in this case, port 3000.

	Next, the kernel marks the corresponding file descriptor for that socket as readable, signaling that data is available. The Go runtime is waiting on this file descriptor via an internal epoll (Linux) or kqueue (BSD/macOS) mechanism. When the runtime is notified, it wakes up the main goroutine, which was previously blocked waiting for new connections in the Accept() call.

	The Go runtime then reads the request data from the socket buffer. Once the connection is accepted, the runtime spawns a new goroutine to handle this specific client request, while the main goroutine goes back to listening for new ones.

	Inside the new goroutine, the HTTP server inspects the request path (e.g., /hello or /about) and dispatches it to the appropriate handler function via the ServeMux router. The handler writes a response, such as “Hello world!” or “I’m Sanny, Junior software engineer,” to the HTTP response writer.

	This response data is written to the socket send buffer, and once it’s ready, the kernel moves the data into the NIC transmit (ring) buffer in RAM. The NIC then converts the binary data back into electromagnetic signals and transmits them over the network to the router, which forwards them to the client. Finally, the client receives the data and displays the web page content.
*/

/*
	HTTP Request-Response Flow in Go Server

	Request Journey (Client → Server):
	The client sends a request through the network. The server's NIC receives it, converts electromagnetic waves to binary, and stores it in RAM. The NIC interrupts the kernel, which copies the data to the socket's receive buffer (checking it's for port 3000). The kernel updates the file descriptor's status to "readable" in the file table (since sockets are treated as files in Unix/Linux). The kernel notifies the Go runtime, which wakes up the main goroutine to accept the connection.

	Processing:
	The Go HTTP server creates a new goroutine to handle the request. This goroutine checks the route (/hello or /about), executes the matching handler, and writes the response to the socket's send buffer.

	Response Journey (Server → Client):
	The kernel copies data from the socket's send buffer to the NIC's transmit buffer. The NIC converts binary to electromagnetic waves and sends it through the network. The router forwards it to the client, which displays the response.
*/

/*
	Visual Flow Diagram
	CLIENT
  │
  │ (1) HTTP Request
  │ (electromagnetic waves)
  ↓
NETWORK (Router/Internet)
  │
  ↓
SERVER - Network Interface Card (NIC)
  │
  │ (2) NIC converts waves → binary data
  │     Stores in NIC Receive Buffer (RAM)
  │
  │ (3) NIC interrupts Kernel
  │
  ↓
KERNEL
  │
  │ (4) Kernel reads from NIC Receive Buffer
  │     Checks port number (3000)
  │     Copies data to Socket Receive Buffer
  │
  │ (4.5) Updates File Descriptor & File Table
  │        - Kernel maintains a file table for all open files/sockets
  │        - Each socket has a file descriptor (integer like 3, 4, 5...)
  │        - Kernel marks this file descriptor as "readable"
  │        - File table entry points to socket structure
  │
  │ (5) Notifies Go Runtime (via file descriptor events)
  │
  ↓
GO RUNTIME
  │
  │ (6) Wakes up main goroutine
  │
  ↓
MAIN GOROUTINE
  │
  │ (7) Calls accept() function
  │     Passes file descriptor to kernel
  │     Requests kernel to read socket
  │
  ↓
KERNEL
  │
  │ (8) Uses file descriptor to locate socket in file table
  │     Reads Socket Receive Buffer
  │     Returns connection data (with new file descriptor for this connection)
  │
  ↓
GO RUNTIME
  │
  │ (9) Creates NEW GOROUTINE to handle request
  │
  ↓
NEW GOROUTINE (Request Handler)
  │
  │ (10) Uses file descriptor for this connection
  │      Reads HTTP request data
  │      Checks request route (/hello or /about)
  │      Executes appropriate handler
  │      Handler writes response
  │
  │ (11) Writes response to Socket Send Buffer (via file descriptor)
  │
  ↓
KERNEL
  │
  │ (12) Notified of data in Send Buffer
  │      Copies data to NIC Transmit Buffer (Ring Buffer)
  │
  ↓
NIC
  │
  │ (13) Reads from Transmit Buffer
  │      Converts binary → electromagnetic waves
  │
  ↓
NETWORK (Router/Internet)
  │
  │ (14) Response travels back
  │
  ↓
CLIENT
  │
  └─→ (15) Displays "Hello world!" or "I'm Sanny..."
*/
