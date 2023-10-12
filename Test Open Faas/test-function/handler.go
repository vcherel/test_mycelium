package function

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var req []byte // Create a variable to store the request body

	// Read input from when the function when it is invoked

	if r.Body != nil {
		// Close the request body when the function exits
		defer r.Body.Close()

		// Read the request body
		body, err := io.ReadAll(r.Body)

		// Check if there was an error reading the request body
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error reading request body: %s", err.Error())))
			return
		}

		req = body
	}

	// Transform input into bytes
	input := string(req)

	// Make an HTTP request to invoke the second function
	resp, err := http.Post("http://test-function-3:8080", "text/plain", strings.NewReader(input))

	// Check if there was an error making the HTTP request
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error making HTTP request: %s", err.Error())))
		return
	}

	// Read input from the second function

	if resp.Body != nil {
		// Close the request body when the function exits
		defer r.Body.Close()

		// Read the request body
		responseBody, err := io.ReadAll(resp.Body)

		// Check if there was an error reading the response body
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error reading response body: %s", err.Error())))
			return
		}

		// Set the HTTP response status code to 200 (OK)
		w.WriteHeader(http.StatusOK)

		// Write the response body
		w.Write([]byte(fmt.Sprintf("%s", string(responseBody))))
	}
}
