package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RunRouter(r *gin.Engine) {
	r.Run(fmt.Sprintf(":%d", Config.Port))
}

func ConfigRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	api := r.Group("/api/v1")

	api.GET("/reports", handleReports)
	api.GET("/sentiments", handleSentiments)
	api.GET("/price", handlePrice)
	api.GET("/member", handleMember)

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204) // Pre-flight request, respond successfully
		} else {
			ctx.Next()
		}
	}
}

func handleReports(c *gin.Context) {
	fname := "reports.json"
	var reports []ReportInfo
	if fileExists(fname) {
		data, err := os.ReadFile(fname)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
			return
		}

		if err := json.Unmarshal(data, &reports); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": reports})
}

func handleSentiments(c *gin.Context) {
	fname := "sentiments.json"
	var sentiments []Sentiment
	if fileExists(fname) {
		data, err := os.ReadFile(fname)
		if err != nil {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
				return
			}
		}

		if err := json.Unmarshal(data, &sentiments); err != nil {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": sentiments})
}

func handlePrice(c *gin.Context) {
	fname := "price_analysis.json"
	var prices []ReportInfo
	if fileExists(fname) {
		data, err := os.ReadFile(fname)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
			return
		}

		if err := json.Unmarshal(data, &prices); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": prices})
}

func handleMember(c *gin.Context) {
	address := c.Query("address")
	fmt.Println("resultgetting for : ", address)

	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Address parameter is required"})
		return
	}

	chainQuery := NewChainQuery()
	isMember := chainQuery.IsMember(address)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "is_member": isMember})
}
