package handler

import (
	"github.com/bastean/laika/pkg/cmd/server/components/layout"
	"github.com/bastean/laika/pkg/cmd/server/components/partials"
	"github.com/gin-gonic/gin"
)

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/")
		layout.Base("home", partials.Button("/dashboard", "Dashboard")).Render(c.Request.Context(), c.Writer)
	}
}
