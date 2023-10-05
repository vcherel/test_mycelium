package function

import (
	"fmt"
	"io"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

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
		input = body
	}

	// Modify the input
	modifiedInput := "Modified: " + string(input)

	// Set the HTTP response status code to 200 (OK)
	w.WriteHeader(http.StatusOK)

	// Write the response body, which includes the content of the 'input' variable, as a string
	w.Write([]byte(fmt.Sprintf("Body: %s", string(modifiedInput))))
}
