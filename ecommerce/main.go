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
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting on the server", err)
	}
}

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
	The client sends a request through the network. The server's NIC receives it, converts electromagnetic waves to binary, and stores it in RAM. The NIC interrupts the kernel, which copies the data to the socket's receive buffer (checking it's for port 3000). The kernel notifies the Go runtime, which wakes up the main goroutine to accept the connection.

	Processing:
	The Go HTTP server creates a new goroutine to handle the request. This goroutine checks the route (/hello or /about), executes the matching handler, and writes the response to the socket's send buffer.

	Response Journey (Server → Client):
	The kernel copies data from the socket's send buffer to the NIC's transmit buffer. The NIC converts binary to electromagnetic waves and sends it through the network. The router forwards it to the client, which displays the response.
*/
