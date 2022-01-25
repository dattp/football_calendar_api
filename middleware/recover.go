package middleware

import (
	"football-calendar-api/common"
	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if errCustom, ok := err.(*common.ErrorCustom); ok {
					c.AbortWithStatusJSON(errCustom.StatusCode, common.ResponseError(errCustom))
					panic(err) // gọi lại panic để recover của gin log ra stack trace error để dễ debug
					return
				}
				errCustom := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(errCustom.StatusCode, common.ResponseError(errCustom))
				panic(err)
				return
			}
		}()
		c.Next()
	}
}
