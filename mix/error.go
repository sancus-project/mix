package mix

import (
	"net/http"

	"go.sancus.dev/mix/types"
)

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {

	// does the error know how to render itself?
	h, ok := err.(http.Handler)
	if !ok {
		var code int

		// but if it doesn't, wrap it in HandlerError{}
		if e, ok := err.(types.Error); ok {
			code = e.Status()
		} else {
			code = http.StatusInternalServerError
		}

		h = &types.HandlerError{
			Code: code,
			Err:  err,
		}
	}

	h.ServeHTTP(w, r)
}
