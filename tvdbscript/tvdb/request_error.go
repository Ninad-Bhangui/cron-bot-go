package tvdb

import "fmt"

type RequestError struct {
    Code int
}

func (e *RequestError) Error() string {
    return fmt.Sprintf("Got response with status code %d", e.Code)
}
