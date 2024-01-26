package services_test

import (
	"testing"

	"github.com/olucasandrade/url-analyzer/services"
	"github.com/stretchr/testify/assert"
)

var service = &services.AnalyzerService{}

func TestErrorWithInvalidURL(t *testing.T) {
	_, err := service.AnalyzeURL("something_invalid")
	assert.Equal(t, "invalid URL format", err.Error())
}

func TestServiceReturnsDataSuccessfully(t *testing.T) {
	response, err := service.AnalyzeURL("criaway.com")
	assert.NotEmpty(t, response.HTTPRequestsCount)
	assert.NotEmpty(t, response.LoadTime)
	assert.NotEmpty(t, response.PageSize)
	assert.Equal(t, nil, err)
}