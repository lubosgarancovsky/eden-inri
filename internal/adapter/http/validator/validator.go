package validator

import (
	"github.com/gin-gonic/gin"
)

func BindAndValidate(c *gin.Context, obj interface{}) error {
	priorityHeader := c.Request.Header.Get("Priority")
	c.Request.Header.Del("Priority")

	if err := c.ShouldBindUri(obj); err != nil {
		return err
	}

	if err := c.ShouldBindHeader(obj); err != nil {
		return err
	}
	if c.Request.Method == "GET" || c.Request.ContentLength == 0 {
		return nil
	}

	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}
	c.Request.Header.Set("Priority", priorityHeader)

	return nil
}
