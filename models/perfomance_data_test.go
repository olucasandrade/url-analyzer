package models_test

import (
	"testing"

	"github.com/olucasandrade/url-analyzer/models"
	"github.com/stretchr/testify/assert"
)

func TestPerfomanceData(t *testing.T) {
	load_time := "3s"
	http_requests_count := 2
	page_size := "10MB"


	perf_data := &models.PerformanceData{
		LoadTime: load_time,
		HTTPRequestsCount: http_requests_count,
		PageSize: page_size,
	}

	assert.Equal(t, load_time, perf_data.LoadTime)
	assert.Equal(t, http_requests_count, perf_data.HTTPRequestsCount)
	assert.Equal(t, page_size, perf_data.PageSize)
}