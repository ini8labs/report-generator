package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/signintech/gopdf"

	"github.com/ini8labs/lsdb"
)

func initializeEventInfo(lottteryEventInfo []lsdb.LotteryEventInfo) []EventsInfo {
	var eventsInfoArr []EventsInfo

	for i := 0; i < len(lottteryEventInfo); i++ {
		eventinfo := EventsInfo{
			EventUID:      primitiveToString(lottteryEventInfo[i].EventUID),
			EventDate:     convertPrimitiveToTime(lottteryEventInfo[i].EventDate),
			EventName:     lottteryEventInfo[i].Name,
			EventType:     lottteryEventInfo[i].EventType,
			WinningNumber: lottteryEventInfo[i].WinningNumber,
		}

		eventsInfoArr = append(eventsInfoArr, eventinfo)
	}
	return eventsInfoArr
}

func (s Server) generateEventInfoPDF(c *gin.Context) {
	resp, err := s.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Error")
	}

	result := initializeEventInfo(resp)

	var eventData []EventsInfo
	b, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to encode JSON data")
		return
	}

	err = json.Unmarshal(b, &eventData)
	if err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeB5,
		Unit: gopdf.UnitCM,
	})
	pdf.AddPage()
	pdf.SetLineWidth(5)

	err = pdf.AddTTFFont("Hamburg", "C:/Users/LENOVO/Desktop/git_offfice/Lottery Project/report-generator/Hamburg.ttf")
	if err != nil {
		s.Logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := pdf.SetFont("Hamburg", "", 10); err != nil {
		s.Logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(eventData); i++ {
		pdf.Cell(nil, "Event ")
		pdf.Cell(nil, (strconv.Itoa(i + 1)))
		pdf.Cell(nil, " :-")
		pdf.Br(0.4)
		pdf.Cell(nil, "Event UID: ")
		pdf.Text(fmt.Sprintf("%v", eventData[i].EventUID))
		pdf.Br(0.3)
		pdf.Cell(nil, "Event Date(DD MM YYYY): ")
		pdf.Text(fmt.Sprintf("%v", eventData[i].EventDate))
		pdf.Br(0.3)
		pdf.Cell(nil, "Event Name: ")
		pdf.Text(fmt.Sprintf("%v", eventData[i].EventName))
		pdf.Br(0.3)
		pdf.Cell(nil, "Event Type: ")
		pdf.Text(fmt.Sprintf("%v", eventData[i].EventType))
		pdf.Br(0.3)
		pdf.Cell(nil, "Winning Numbers: ")
		pdf.Text(fmt.Sprintf("%v", eventData[i].WinningNumber))
		pdf.Br(0.5)

	}

	fileName := "eventinfo.pdf"
	filePath := fmt.Sprintf("C:/Users/LENOVO/Desktop/check-point/%s", fileName)

	err = pdf.WritePdf(filePath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "PDF generated and stored successfully")

}

func (s Server) downloadLink(c *gin.Context) {
	fileName := "eventinfo.pdf"
	filePath := fmt.Sprintf("C:/Users/LENOVO/Desktop/check-point/%s", fileName)

	// Check if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Generate a download link for the PDF
	downloadLink := fmt.Sprintf("http://%s/download-pdf?filePath=%s", c.Request.Host, filePath)

	c.JSON(http.StatusOK, gin.H{"downloadLink": downloadLink})
}

func (s Server) downloadPDF(c *gin.Context) {
	filePath := c.Query("filePath")

	// Check if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Set the response headers for PDF download
	c.Header("Content-Disposition", "attachment; filename=example.pdf")
	c.Header("Content-Type", "application/pdf")

	// Stream the file directly to the response
	c.File(filePath)
}
