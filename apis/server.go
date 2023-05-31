package apis

import (
	"github.com/gin-gonic/gin"
)

func NewServer(server Server) error {

	r := gin.Default()

	// API end point\
	r.GET("/api/v1/generate-report/events", server.generateEventReport)
	r.GET("/api/v1/generate-report/winners", server.generateWinnersReport)
	//r.POST("/api/v1/download-pdf", server.downloadPDF)

	return r.Run(server.Addr)
}
