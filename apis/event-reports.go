package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"

	apis "github.com/ini8labs/admin-service/src/apis"
)

const (
	bannerHt = 43
	xIndent  = 5
	yIndent  = 47
)

func drawFooter(pdf *gofpdf.Fpdf, w float64, h float64, bannerHt float64) {
	pdf.SetFillColor(103, 60, 79)
	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: h},
		{X: 0, Y: h - 15},
		{X: w, Y: h - 15},
		{X: w, Y: h},
	}, "F ")
}

func drawBanner(pdf *gofpdf.Fpdf, w float64, h float64, bannerHt float64) {

	pdf.SetFillColor(103, 60, 79)

	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: 0},
		{X: w, Y: 0},
		{X: w, Y: bannerHt},
		{X: 0, Y: bannerHt},
	}, "F ")
	pdf.SetFont("Arial", "B", 50)
	pdf.SetTextColor(255, 255, 255)
	pdf.Text(5, 21, "Lottery Project")
	pdf.SetFont("Arial", "I", 25)
	pdf.Text(10, 38, "Events Info")
}

// func drawGrid(pdf *gofpdf.Fpdf, w float64, h float64) {
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.SetTextColor(255, 0, 0)
// 	pdf.SetDrawColor(200, 200, 200)

//		for x := 0.0; x < w; x = x + (w / 20) {
//			pdf.Line(x, 0, x, h)
//			_, lineht := pdf.GetFontSize()
//			pdf.Text(x, lineht, fmt.Sprintf("%d", int(x)))
//		}
//		for y := 0.0; y < h; y = y + (w / 20) {
//			pdf.Line(0, y, w, y)
//			pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
//		}
//	}

func drawColumnAndInsertHeaders(pdf *gofpdf.Fpdf, x float64, y float64, pageNuM float64) {

	height := 11.0

	pdf.SetFont("times", "", 15)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(211, 211, 211)

	pdf.MoveTo(x, y)
	pdf.CellFormat(15, height, "S.No", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+15, y)
	pdf.CellFormat(30, height, "Name", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+35+10, y)
	pdf.CellFormat(45, height, "Date", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+(35+10+45), y)
	pdf.CellFormat(50, height, "UID", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+(35+10+45+50), y)
	pdf.CellFormat(25, height, "Type", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+(35+10+45+50+25), y)
	pdf.CellFormat(42, height, "Winning Numbers", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")
}

func insertValues(pdf *gofpdf.Fpdf, eventsInfo []apis.EventsInfo, x float64, y float64) {
	pdf.SetFont("times", "", 11)
	pdf.SetTextColor(0, 0, 0)
	breadth := 11.0
	pageNum := 2.0

	for i := 0; i < len(eventsInfo); i++ {
		height := 11.0 * (float64(i) + 1)

		if i < 18 {

			pdf.MoveTo(x, y+height)
			pdf.CellFormat(15, breadth, strconv.Itoa(i+1), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15, y+height)
			pdf.CellFormat(30, breadth, eventsInfo[i].EventName, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10, y+height)
			pdf.CellFormat(45, breadth, convertDateToString(eventsInfo[i].EventDate), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45, y+height)
			pdf.CellFormat(50, breadth, eventsInfo[i].EventUID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45+50, y+height)
			pdf.CellFormat(25, breadth, eventsInfo[i].EventType, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45+50+25, y+height)
			pdf.CellFormat(42, breadth, convertWinNumbersToString(eventsInfo[i].WinningNumber), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")
		}

		if (i == 18) || (i == 36) || (i == 54) || (i == 72) || (i == 90) {

			pdf.AddPage()
			w, h := pdf.GetPageSize()
			drawBanner(pdf, w, h, bannerHt)
			drawFooter(pdf, w, h, bannerHt)
			drawColumnAndInsertHeaders(pdf, xIndent, yIndent, pageNum)
			pdf.SetFont("times", "", 11)
			pdf.SetTextColor(0, 0, 0)

			pdf.MoveTo(x, y+(breadth))
			pdf.CellFormat(15, breadth, strconv.Itoa(i+1), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15, y+breadth)
			pdf.CellFormat(30, breadth, eventsInfo[i].EventName, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10, y+breadth)
			pdf.CellFormat(45, breadth, convertDateToString(eventsInfo[i].EventDate), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45, y+breadth)
			pdf.CellFormat(50, breadth, eventsInfo[i].EventUID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45+50, y+breadth)
			pdf.CellFormat(25, breadth, eventsInfo[i].EventType, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45+50+25, y+breadth)
			pdf.CellFormat(42, breadth, convertWinNumbersToString(eventsInfo[i].WinningNumber), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")
			pageNum++
		}

		if (i > 18 && i < 36) || (i > 36 && i < 54) || (i > 54 && i < 72) || (i > 72 && i < 90) {
			q := 2.0
			pdf.MoveTo(x, y+(breadth*q))
			pdf.CellFormat(15, breadth, strconv.Itoa(i+1), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15, y+breadth*q)
			pdf.CellFormat(30, breadth, eventsInfo[i].EventName, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10, y+breadth*q)
			pdf.CellFormat(45, breadth, convertDateToString(eventsInfo[i].EventDate), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45, y+breadth*q)
			pdf.CellFormat(50, breadth, eventsInfo[i].EventUID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45+50, y+breadth*q)
			pdf.CellFormat(25, breadth, eventsInfo[i].EventType, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45+50+25, y+breadth*q)
			pdf.CellFormat(42, breadth, convertWinNumbersToString(eventsInfo[i].WinningNumber), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")
			q++
		}
	}
}

func (s Server) generateEventReport(c *gin.Context) {

	resp, err := s.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Error")
	}

	result := apis.InitializeEventInfo(resp)
	fmt.Println(result)

	var eventData []apis.EventsInfo
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

	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.AddPage()
	pdf.MoveTo(0, 0)
	pdf.SetFont("Arial", "B", 30)

	w, h := pdf.GetPageSize()

	drawBanner(pdf, w, h, bannerHt)
	drawFooter(pdf, w, h, bannerHt)
	drawColumnAndInsertHeaders(pdf, xIndent, yIndent, 1)
	insertValues(pdf, result, xIndent, yIndent)

	//pdf.SetFillColor(103, 60, 79)

	//drawGrid(pdf, w, h)

	err = pdf.OutputFileAndClose("event-info.pdf")
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}

}
