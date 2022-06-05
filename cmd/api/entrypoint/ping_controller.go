package entrypoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

// Ping returns a successful pong answer to all HTTP requests
func (controller *PingController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
