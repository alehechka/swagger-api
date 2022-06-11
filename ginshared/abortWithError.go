package ginshared

import (
	"net/http"

	"github.com/alehechka/swagger-api/types"
	"github.com/gin-gonic/gin"
)

// ShouldAbortWithError takes a status and array of errors and will abort if any errors are non-nil.
// Bool return value indicates whether or not to short-circuit
func ShouldAbortWithError(c *gin.Context) func(status int, errs ...error) bool {
	return func(status int, errs ...error) bool {
		var errors types.Errors

		for _, err := range errs {
			if err != nil {
				c.Error(err)
				errors = append(errors, types.Error{
					Status: status,
					Title:  http.StatusText(status),
					Detail: err.Error(),
				})
			}
		}

		if len(errors) > 0 {
			c.AbortWithStatusJSON(status, types.ErrorsResponse{
				Errors: errors,
			})
			return true
		}

		c.Next()
		return false
	}
}
