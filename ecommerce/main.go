package main

import (
	"ecommerce/cmd"
)

func main() {
	cmd.Serve()
}

/*
	[] = list
	{} = object
	json = javascript object notation

	If struct field names start with lowercase (e.g. id), they’re unexported and won’t appear in JSON output.

	Keep them uppercase (e.g. ID) and use JSON tags to control key names, like:

	type Product struct {
		ID int `json:"id"`
	}

	The encoding/json package is a separate package, so it can’t access unexported (lowercase) struct fields because they’re private to your package.

	So — lowercase fields = private → invisible to json.
	Uppercase fields = exported → visible to json.

	CORS = cross origin resource sharing
	It’s a security feature in web browsers that controls which websites can request data from your server.


	OPTIONS is an HTTP method.
 	preflightReq = Browser automatically sends a preflight request (using the OPTIONS method) before making certain cross-origin requests. This request checks if the server allows it by verifying the CORS headers in the response.

	SOLID's S is single responsibility principle that means one func should do one particular job, we'll discover more of principle later.

*/
