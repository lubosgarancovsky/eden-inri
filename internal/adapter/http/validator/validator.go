package validator

import "github.com/gin-gonic/gin"

func BindAndValidate(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindHeader(obj); err != nil {
		return err
	}

	if err := c.ShouldBindUri(obj); err != nil {
		return err
	}

	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}

	return nil
}
