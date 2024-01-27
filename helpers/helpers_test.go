package helpers_test

import (
	"testing"

	"github.com/olucasandrade/url-analyzer/helpers"
	"github.com/stretchr/testify/assert"
)

func TestHasHTTPPrefix(t *testing.T) {
	url_with_http := "http://www.criaway.com"
	url_without_http := "www.criaway.com"
	empty_string := ""

	assert.Equal(t, true, helpers.HasHTTPPrefix(url_with_http))
	assert.Equal(t, false, helpers.HasHTTPPrefix(url_without_http),)
	assert.Equal(t, false, helpers.HasHTTPPrefix(empty_string))
}

func TestIsValidURL(t *testing.T) {
	valid_url := "http://www.criaway.com"
	invalid_url_with_http := "http://something..."
	url_without_http := "criaway.com"
	email := "lafdesouza2002@gmail.com"
	empty_string := ""
	special_characters := "!@#$ˆ&*()"

	assert.Equal(t, true, helpers.IsValidURL(valid_url))
	assert.Equal(t, false, helpers.IsValidURL(invalid_url_with_http))
	assert.Equal(t, false, helpers.IsValidURL(url_without_http))
	assert.Equal(t, false, helpers.IsValidURL(email))
	assert.Equal(t, false, helpers.IsValidURL(empty_string))
	assert.Equal(t, false, helpers.IsValidURL(special_characters))
}