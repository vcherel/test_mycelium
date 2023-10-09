package function

import (
	"fmt"
	"io"
	"net/http"
)

// Read input from the request body and write the response body with the modified input
func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error reading request body: %s", err.Error())))
			return
		}

		input = body
	}

	modifiedInput := "Modified: " + string(input)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s", string(modifiedInput))))
}
