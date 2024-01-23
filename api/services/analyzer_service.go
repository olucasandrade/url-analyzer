package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/olucasandrade/url-analyzer/api/models"
)

func convertPageSizeToMB(pageSize int) string {
	return fmt.Sprintf("%.4fMB", float64(pageSize)/(1024*1024))
}

func hasHTTPPrefix(url string) bool {
	return len(url) >= 7 && (url[:7] == "http://" || url[:8] == "https://")
}

type AnalyzerService struct {}

func (as *AnalyzerService) AnalyzeURL(url string) (*models.PerformanceData, error) {
	if !hasHTTPPrefix(url) {
		url = "http://" + url
	}

	startTime := time.Now()

	if startTime.IsZero() {
		return nil, errors.New("failed to initialize startTime")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	loadTime := fmt.Sprintf("%.2fs", time.Since(startTime).Seconds())

	httpRequestsCount := doc.Find("a[href], img, script, link[rel='stylesheet']").Length()

	pageSize := convertPageSizeToMB(len([]byte(doc.Text())))

	return &models.PerformanceData{
		LoadTime:          loadTime,
		HTTPRequestsCount: httpRequestsCount,
		PageSize:          pageSize,
	}, nil
}