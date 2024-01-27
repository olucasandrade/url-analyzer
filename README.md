# URL Analyzer

URL Analyzer is a Go package for analyzing website performance metrics.

## Installation

To install the package, use the `go get` command:

```bash
go get -u github.com/olucasandrade/url-analyzer
```

## Usage

Here's a basic example of how to use `url-analyzer` package:

```go
package main

import (
	"fmt"
    "github.com/olucasandrade/url-analyzer"
)

func main() {
	response, err := url_analyzer.Analyze("www.criaway.com")
	if (err != nil) {
		fmt.Println(err)
	} else {
        fmt.Println(response.LoadTime)
    }
}

```

## Types

`PerformanceData` represents the performance metrics obtained after analyzing a URL.

```go
type PerformanceData struct {
    LoadTime           string // Time taken for the page to load
    HTTPRequestsCount  int           // Number of HTTP requests made during page load
    PageSize           int           // Size of the page in bytes
}
```

## Testing

To run tests for the url-analyzer package, use the following command:

```bash
go test -v ./...
```

## Contributions

Contributions are welcome! If you find any issues or have ideas to improve the package, feel free to open an issue or submit a pull request.

## License

This package is distributed under the MIT license. See the LICENSE file for more information.