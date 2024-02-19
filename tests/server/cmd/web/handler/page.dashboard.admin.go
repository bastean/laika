package handler

import (
	"github.com/bastean/laika/test/server/cmd/web/components/layout"
	"github.com/bastean/laika/test/server/cmd/web/components/page"
	"github.com/gin-gonic/gin"
)

func IndexDashboardAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/dashboard/admin")
		layout.Base("admin", page.Admin()).Render(c.Request.Context(), c.Writer)
	}
}
