package apis

import (
	"github.com/gin-gonic/gin"
)

func NewServer(server Server) error {

	r := gin.Default()

	// API end point\
	r.GET("/api/v1/generate-pdf/events", server.generateEventInfoPDF)
	r.GET("/api/v1/download-link", server.downloadLink)
	r.POST("/api/v1/download-pdf", server.downloadPDF)

	return r.Run(server.Addr)
}
