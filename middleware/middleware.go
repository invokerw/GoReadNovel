package middleware

import (
	"GoReadNovel/logger"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	logger.ALogger().Debug("   ")
}
