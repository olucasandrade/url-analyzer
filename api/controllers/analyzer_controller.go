package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olucasandrade/url-analyzer/api/services"
)

type AnalyzeController struct {
    AnalyzerService services.AnalyzerService
}

type Payload struct {
	Url string `json:"url"`
}

func (ac *AnalyzeController) AnalyzeHandler(ctx *gin.Context)() {
	var payload Payload

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "please send a valid body"})
		return
	}

	url := payload.Url

	if url == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL cannot be empty"})
		return
	}

	performanceData, err := ac.AnalyzerService.AnalyzeURL(url)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error while analyzing data"})
	}

	ctx.JSON(200, gin.H{
		"url":                 url,
		"load_time":           performanceData.LoadTime,
		"http_requests_count": performanceData.HTTPRequestsCount,
		"page_size":           performanceData.PageSize,
	})
}