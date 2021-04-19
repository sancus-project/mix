package mix

import (
	"net/http"
)

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	panic(err)
}
