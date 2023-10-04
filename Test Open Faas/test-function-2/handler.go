package function

import (
	"fmt"
	"io"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var req []byte

	// Check if the request body is not nil
	if r.Body != nil {
		// Close the request body when the function exits
		defer r.Body.Close()

		// Read the entire request body into the 'body' variable
		body, _ := io.ReadAll(r.Body)

		// Assign the 'body' variable (as bytes) to the 'input' variable
		req = body
	}

	// Modify the input
	modifiedInput := "Modified: " + string(req)

	// Print the modified input to the console
	fmt.Println(modifiedInput)

	// Set the HTTP response status code to 200 (OK)
	w.WriteHeader(http.StatusOK)

	// Write the response body, which includes the content of the 'input' variable, as a string
	w.Write([]byte(fmt.Sprintf("Body: %s", string(modifiedInput))))
}