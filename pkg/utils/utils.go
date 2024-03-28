package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseRequestBody(r *http.Request, x interface{}) {
	// Read the body of the request.
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
