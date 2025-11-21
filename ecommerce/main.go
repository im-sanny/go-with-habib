package main

import (
	"ecommerce/cmd"
)

func main() {

	cmd.Serve()
	// jwt, err := util.CreateJwt("my_secret", util.Payload{
	// 	Sub:         45,
	// 	FirsName:    "Theo",
	// 	LastName:    "Fumis",
	// 	Email:       "hello@gmail.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(jwt)
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

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Name") // for this to work i need to set a custom header from frontend code

preflightReq = Browser automatically sends a preflight request (using the OPTIONS method) before making certain cross-origin requests. This request checks if the server allows it by verifying the CORS headers in the response.

SOLID's S is single responsibility principle that means one func should do one particular job, we'll discover more of principle later.

Base64: A way to encode binary data (like images, files, or bytes) into text using only 64 safe ASCII characters (A-Z, a-z, 0-9, +, /, =). By using this we can safely send binary data over systems that only handle text — like emails, JSON, URLs, or APIs. It’s for encoding, not security. Anyone can decode it easily.

SHA-256: A cryptographic hash function that takes any input (text, file, etc.) and produces a unique 256-bit (64-character hexadecimal) fingerprint. Same input → always same hash. Tiny change in input → totally different hash. One-way: You cannot reverse it to get the original data. Fast to compute. Can be used Verifying file integrity (e.g., “Did this download get corrupted?”), Storing passwords (with salt!), Blockchain, digital signatures.

HMAC: Hash-based Message Authentication Code — a way to verify that a message hasn’t been tampered with AND came from someone who knows a secret key. It combines a secret key + your message + a hash function (like SHA-256) to produce a code. It gets used to prove authenticity — not just that data is unchanged, but that it came from a trusted source.

HMAC-SHA256: The most common real-world version of HMAC — it uses SHA-256 as the hash function. It get's used Securing API requests (e.g., AWS, GitHub webhooks). Generating secure tokens. Authenticating users in stateless systems. It's popular because it's Fast, secure, widely supported, and resistant to attacks.

Formula:
HMAC-SHA256(secret_key, message) → outputs a 64-character hex string

JWT (JSON Web Token): A standardized format for securely transmitting information between parties as a JSON object.
Commonly used for authentication and session management in web apps. JWT has 3 parts as it's structure header.payload.signature.

*/
