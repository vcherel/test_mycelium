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
		body, err := io.ReadAll(r.Body)

		// Check if there was an error reading the request body
		if err != nil {
			// Set the HTTP response status code to 500 (Internal Server Error)
			w.WriteHeader(http.StatusInternalServerError)
			
			// Write the error message as the response body
			w.Write([]byte(fmt.Sprintf("Error reading request body: %s", err.Error())))

			// Exit the function
			return
		}

		// Assign the 'body' variable (as bytes) to the 'input' variable
		req = body
	}

	// Read input from the request body
	input := string(req)

	// Make an HTTP request to invoke the second function
	resp, err := http.Post("http://127.0.0.1:8080/function/test-function-2", "application/json", strings.NewReader(input))

	// Check if there was an error making the HTTP request
	if err != nil {
		// Set the HTTP response status code to 500 (Internal Server Error)
		w.WriteHeader(http.StatusInternalServerError)

		// Write the error message as the response body
		w.Write([]byte(fmt.Sprintf("Error making HTTP request: %s", err.Error())))

		// Exit the function
		return
	}

	// Check if the response body is not nil
	if resp.Body != nil {
		// Read the response from the second function
		responseBody, err := io.ReadAll(resp.Body)

		// Check if there was an error reading the response body
		if err != nil {
			// Set the HTTP response status code to 500 (Internal Server Error)
			w.WriteHeader(http.StatusInternalServerError)

			// Write the error message as the response body
			w.Write([]byte(fmt.Sprintf("Error reading response body: %s", err.Error())))

			// Exit the function
			return
		}

		// Set the HTTP response status code to 200 (OK)
		w.WriteHeader(http.StatusOK)

		// Write the response body, which includes the content of the 'input' variable, as a string
		w.Write([]byte(fmt.Sprintf("Body: %s", string(responseBody))))
	}
}
