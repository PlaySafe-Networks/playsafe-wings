package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// postServerHibernate is a Phase 0 placeholder for the PlaySafe hibernation
// subsystem. The full implementation arrives in Phase 4 and performs a
// graceful docker stop, persists hibernation state, and emits a "hibernated"
// event back to the Panel.
func postServerHibernate(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "hibernation not implemented in this Wings build",
	})
}

// postServerWake is a Phase 0 placeholder for the PlaySafe hibernation
// subsystem. The full implementation arrives in Phase 4 and performs a
// docker start, polls the container healthcheck, and emits a "running" event
// back to the Panel when the server is ready.
func postServerWake(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "wake not implemented in this Wings build",
	})
}
