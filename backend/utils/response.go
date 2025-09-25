// utils/response.go
package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

type HTTPResponder struct{}

func NewHTTPResponder() *HTTPResponder {
	return &HTTPResponder{}
}

func (r *HTTPResponder) Success(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func (r *HTTPResponder) SuccessWithMessage(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func (r *HTTPResponder) Error(c *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
	c.JSON(statusCode, gin.H{"error": message})
}

func (r *HTTPResponder) BadRequest(c *gin.Context, message string, err error) {
	r.Error(c, 400, message, err)
}

func (r *HTTPResponder) InternalError(c *gin.Context, message string, err error) {
	r.Error(c, 500, message, err)
}

func (r *HTTPResponder) Unauthorized(c *gin.Context, message string) {
	r.Error(c, 401, message, nil)
}

func (r *HTTPResponder) NotFound(c *gin.Context, message string) {
	r.Error(c, 404, message, nil)
}
