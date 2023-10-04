package function

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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

	// Read input from the request body
	input := string(req)

	// Make an HTTP request to invoke the second function
	resp, _ := http.Post("http://127.0.0.1:8080/function/test-function-2", "application/json", strings.NewReader(input))

	// Check if the response body is not nil
	if resp.Body != nil {
		// Read the response from the second function
		responseBody, _ := io.ReadAll(resp.Body)

		// Set the HTTP response status code to 200 (OK)
		w.WriteHeader(http.StatusOK)

		// Write the response body, which includes the content of the 'input' variable, as a string
		w.Write([]byte(fmt.Sprintf("Body: %s", string(responseBody))))
	}
}
