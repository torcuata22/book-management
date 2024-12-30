package utils

import (
	"encoding/json"
	"io"

	//"os"
	"net/http"
)

func ParseBody(r *http.Request, result interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), result); err != nil {
			panic(err)
		}
	}

}
