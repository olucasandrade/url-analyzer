package url_analyzer_test

import (
	"testing"

	url_analyzer "github.com/olucasandrade/url-analyzer"
	"github.com/stretchr/testify/assert"
)

func TestServiceReturnsDataSuccessfully(t *testing.T) {
	response, err := url_analyzer.Analyze("criaway.com")
	assert.NotEmpty(t, response.HTTPRequestsCount)
	assert.NotEmpty(t, response.LoadTime)
	assert.NotEmpty(t, response.PageSize)
	assert.Equal(t, nil, err)
}