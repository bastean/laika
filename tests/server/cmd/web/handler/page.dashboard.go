package handler

import (
	"github.com/bastean/laika/test/server/cmd/web/components/layout"
	"github.com/bastean/laika/test/server/cmd/web/components/partials"
	"github.com/gin-gonic/gin"
)

func IndexDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/dashboard")
		layout.Base("dashboard", partials.Button("/dashboard/admin", "Admin")).Render(c.Request.Context(), c.Writer)
	}
}
