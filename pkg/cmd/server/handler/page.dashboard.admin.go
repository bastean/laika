package handler

import (
	"github.com/bastean/laika/pkg/cmd/server/components/layout"
	"github.com/bastean/laika/pkg/cmd/server/components/partials"
	"github.com/gin-gonic/gin"
)

func IndexDashboardAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/dashboard/admin")
		layout.Base("admin", partials.Button("/", "Home")).Render(c.Request.Context(), c.Writer)
	}
}
