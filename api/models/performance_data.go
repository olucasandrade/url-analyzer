package models

type PerformanceData struct {
    LoadTime           string `json:"loadTime"`
    HTTPRequestsCount  int           `json:"httpRequestsCount"`
    PageSize           string           `json:"pageSize"`
}