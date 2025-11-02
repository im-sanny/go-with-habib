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
	clint request goes to server-> server send request to NIC -> NIC convert electromagnetic wave into binary and this data will keep in RAM -> in ram there is NIC receive buffer and write buffer -> now NIC will interrupt karnel and when it interrupt the karnel, karnel will understand that he got request then it'll read data from the RAM then it'll go to another buffer in RAM which is for socket, it'll read that data and see in which port this data came and it'll see that data came in 3000 port, now it'll keep data in  allocated received buffer of the socket after that it'll go to the file descriptor then for sockets file descriptor will mark the file for socket as readable, now karnel will send this information to go runtime, now go routine will find out which data/information it gave, now go runtime will tell main goroutine and awake it from sleep and it sees the readable data and then it want to read that data and tell go runtime that it want's to read that data then goroutine  will tell karnel that it wants to read that data and karnel thinks this request send from the go process not from goroutine so then sockets receive buffer will be read by main goroutine, after this the data it took from socket will read and return from accept function inside main goroutine then it'll send this data to go runtime to handle, and go runtime will create a new goroutine to handle the data, and when it comes to the new goroutine to handle the data is sees that what route mux has, and it sees for which routes request came and by request wise it executes handler it gets, and now it writes in sockets send buffer, after data writes in send buffer sockets interrupt karnel and karnel understand it's ready then karnel writes data in RAM's ring buffer, and now NIC will send data from ring buffer as electro magnetic wave to router and then router sends the data to client and then client shows the data it requested for.
*/


/*
	When a client sends an HTTP request to the server, the request first travels through the network and reaches the server’s Network Interface Card (NIC). The NIC converts the incoming electromagnetic signals into binary data and stores it temporarily in the RAM, specifically inside the NIC receive buffer.

	After this, the NIC triggers an interrupt to notify the kernel that new data has arrived. The kernel then reads the data from the NIC buffer and moves it to another area in RAM called the socket receive buffer, which is associated with the port the server is listening on — in this case, port 3000.

	Next, the kernel marks the corresponding file descriptor for that socket as readable, signaling that data is available. The Go runtime is waiting on this file descriptor via an internal epoll (Linux) or kqueue (BSD/macOS) mechanism. When the runtime is notified, it wakes up the main goroutine, which was previously blocked waiting for new connections in the Accept() call.

	The Go runtime then reads the request data from the socket buffer. Once the connection is accepted, the runtime spawns a new goroutine to handle this specific client request, while the main goroutine goes back to listening for new ones.

	Inside the new goroutine, the HTTP server inspects the request path (e.g., /hello or /about) and dispatches it to the appropriate handler function via the ServeMux router. The handler writes a response, such as “Hello world!” or “I’m Sanny, Junior software engineer,” to the HTTP response writer.

	This response data is written to the socket send buffer, and once it’s ready, the kernel moves the data into the NIC transmit (ring) buffer in RAM. The NIC then converts the binary data back into electromagnetic signals and transmits them over the network to the router, which forwards them to the client. Finally, the client receives the data and displays the web page content.
*/
