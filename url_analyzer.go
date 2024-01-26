package urlanalyzer

import (
	"github.com/olucasandrade/url-analyzer/models"
	"github.com/olucasandrade/url-analyzer/services"
)

func Analyze(url string) (*models.PerformanceData, error) {
    as := &services.AnalyzerService{}
    return as.AnalyzeURL(url)
}