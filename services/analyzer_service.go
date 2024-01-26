package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/olucasandrade/url-analyzer/helpers"
	"github.com/olucasandrade/url-analyzer/models"
)

type AnalyzerService struct {}

func (as *AnalyzerService) AnalyzeURL(url string) (*models.PerformanceData, error) {
	if !helpers.HasHTTPPrefix(url) {
		url = "http://" + url
	}

	if !helpers.IsValidURL(url) {
		return nil, errors.New("invalid URL format")
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

	pageSize := helpers.ConvertPageSizeToMB(len([]byte(doc.Text())))

	return &models.PerformanceData{
		LoadTime:          loadTime,
		HTTPRequestsCount: httpRequestsCount,
		PageSize:          pageSize,
	}, nil
}