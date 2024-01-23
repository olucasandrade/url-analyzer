package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olucasandrade/url-analyzer/api/controllers"
)

func SetupRoutes(ac *controllers.AnalyzeController) *gin.Engine {
	r := gin.Default()
	r.POST("/analyze", ac.AnalyzeHandler)
	return r
}