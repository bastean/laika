package handler

import (
	"github.com/bastean/laika/server/components/layout"
	"github.com/bastean/laika/server/components/partials"
	"github.com/gin-gonic/gin"
)

func IndexDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/dashboard")
		layout.Base("dashboard", partials.Button("/dashboard/admin", "Admin")).Render(c.Request.Context(), c.Writer)
	}
}
