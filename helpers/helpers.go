package helpers

import (
	"fmt"
	"regexp"
)

func ConvertPageSizeToMB(pageSize int) string {
	return fmt.Sprintf("%.4fMB", float64(pageSize)/(1024*1024))
}

func HasHTTPPrefix(url string) bool {
	return len(url) >= 7 && (url[:7] == "http://" || url[:8] == "https://")
}

func IsValidURL(url string) bool {
	matched, err := regexp.MatchString(`^(http|https)://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?$`, url)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return matched
}