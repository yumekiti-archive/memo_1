package router

import (
	"backend/interface/handlers"
	"github.com/gin-gonic/gin"
)

// routes
func NewRouter(
	r *gin.Engine,
	memoHandler handlers.MemoHandler,
) {
	// /api group
	api := r.Group("/api")

	// memo routes
	api.POST("/memos", memoHandler.Create)
	api.GET("/memos/:hash", memoHandler.Read)
	api.GET("/memos", memoHandler.ReadAll)
	api.PUT("/memos/:hash", memoHandler.Update)
	api.DELETE("/memos/:hash", memoHandler.Delete)

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
}
