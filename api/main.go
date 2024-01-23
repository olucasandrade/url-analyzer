package main

import (
	"github.com/olucasandrade/url-analyzer/api/controllers"
	"github.com/olucasandrade/url-analyzer/api/routes"
	"github.com/olucasandrade/url-analyzer/api/services"
)

func main() {
    analyzerService := &services.AnalyzerService{}
    analyzeController := &controllers.AnalyzeController{AnalyzerService: *analyzerService}

	g := routes.SetupRoutes(analyzeController)
	g.Run(":5050")
}
