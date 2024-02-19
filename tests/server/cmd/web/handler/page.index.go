package handler

import (
	"github.com/bastean/laika/test/server/cmd/web/components/layout"
	"github.com/bastean/laika/test/server/cmd/web/components/page"
	"github.com/gin-gonic/gin"
)

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/")
		layout.Base("laika", page.Home()).Render(c.Request.Context(), c.Writer)
	}
}
