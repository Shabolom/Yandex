package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestBinderBody проверяет валидность пришедшего body
func RequestBinderBody(model interface{}, c *gin.Context) error {
	if err := c.ShouldBind(model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return err
	}
	return nil
}

// RequestBinderURI проверяет валидность пришедшего URI
func RequestBinderURI(model interface{}, c *gin.Context) error {
	if err := c.ShouldBindUri(model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return err
	}
	return nil
}
