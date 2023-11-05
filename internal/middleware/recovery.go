package middleware

import (
	"github.com/HarryLuo227/simple-blog-service/global"
	"github.com/HarryLuo227/simple-blog-service/pkg/app"
	"github.com/HarryLuo227/simple-blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrames().Errorf(c, s, err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
