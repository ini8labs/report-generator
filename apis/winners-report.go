package apis

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ini8labs/lsdb"
	"github.com/jung-kurt/gofpdf"
)

type Winners struct {
	UserID    string `json:"user_id"`
	EventUID  string `json:"event_id"`
	AmountWon int    `json:"amountWon"`
	WinType   string `json:"winType"`
}

func InitializeWinnersInfo(eventWinnerInfo []lsdb.WinnerInfo, eventParticipantInfo []lsdb.EventParticipantInfo) []Winners {
	var winnerInfoArr []Winners

	for i := 0; i < len(eventWinnerInfo); i++ {
		winnerInfo := Winners{
			EventUID:  primitiveToString(eventWinnerInfo[i].EventID),
			UserID:    primitiveToString(eventWinnerInfo[i].UserID),
			AmountWon: eventWinnerInfo[i].AmountWon,
			WinType:   eventWinnerInfo[i].WinType,
		}
		winnerInfoArr = append(winnerInfoArr, winnerInfo)
	}

	return winnerInfoArr
}

func drawColumnAndInsertWinnersHeaders(pdf *gofpdf.Fpdf, x float64, y float64, pageNum float64) {
	height := 11.0

	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(211, 211, 211)

	pdf.MoveTo(x+5, y)
	pdf.CellFormat(15, height, "S.No", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+15+5, y)
	pdf.CellFormat(50, height, "User ID", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+15+5+50, y)
	pdf.CellFormat(50, height, "Event Id", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+(15+5+50+50), y)
	pdf.CellFormat(40, height, "Win Type", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

	pdf.MoveTo(x+(15+5+50+50+40), y)
	pdf.CellFormat(40, height, "Amount Won", gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, true, 0, "")

}

func insertWinnerInfo(pdf *gofpdf.Fpdf, winnerInfo []Winners, x float64, y float64) {
	pdf.SetFont("times", "", 11)
	pdf.SetTextColor(0, 0, 0)
	breadth := 11.0
	pageNum := 2.0

	for i := 0; i < len(winnerInfo); i++ {
		height := 11.0 * (float64(i) + 1)

		if i < 18 {
			pdf.MoveTo(x+5, y+height)
			pdf.CellFormat(15, breadth, strconv.Itoa(i+1), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15+5, y+height)
			pdf.CellFormat(50, breadth, winnerInfo[i].UserID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15+5+50, y+height)
			pdf.CellFormat(50, breadth, winnerInfo[i].EventUID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15+5+50+50, y+height)
			pdf.CellFormat(40, breadth, winnerInfo[i].WinType, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15+5+50+50+40, y+height)
			pdf.CellFormat(40, breadth, strconv.Itoa(winnerInfo[i].AmountWon), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

		}
		if (i == 18) || (i == 36) || (i == 54) || (i == 72) || (i == 90) {

			pdf.AddPage()
			w, h := pdf.GetPageSize()
			drawBanner(pdf, w, h, bannerHt, "Event Winners")
			drawFooter(pdf, w, h, bannerHt)
			drawColumnAndInsertEventHeaders(pdf, xIndent, yIndent, pageNum)
			pdf.SetFont("times", "", 11)
			pdf.SetTextColor(0, 0, 0)

			pdf.MoveTo(x, y+(breadth))
			pdf.CellFormat(15, breadth, strconv.Itoa(i+1), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15, y+breadth)
			pdf.CellFormat(30, breadth, winnerInfo[i].UserID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10, y+breadth)
			pdf.CellFormat(45, breadth, winnerInfo[i].EventUID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45, y+breadth)
			pdf.CellFormat(50, breadth, winnerInfo[i].EventUID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45+50, y+breadth)
			pdf.CellFormat(25, breadth, strconv.Itoa(winnerInfo[i].AmountWon), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pageNum++
		}
		if (i > 18 && i < 36) || (i > 36 && i < 54) || (i > 54 && i < 72) || (i > 72 && i < 90) {
			q := 2.0
			pdf.MoveTo(x, y+(breadth*q))
			pdf.CellFormat(15, breadth, strconv.Itoa(i+1), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+15, y+breadth*q)
			pdf.CellFormat(30, breadth, winnerInfo[i].UserID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10, y+breadth*q)
			pdf.CellFormat(45, breadth, winnerInfo[i].EventUID, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45, y+breadth*q)
			pdf.CellFormat(50, breadth, winnerInfo[i].WinType, gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			pdf.MoveTo(x+35+10+45+50, y+breadth*q)
			pdf.CellFormat(25, breadth, strconv.Itoa(winnerInfo[i].AmountWon), gofpdf.BorderFull, gofpdf.LineBreakNone, gofpdf.AlignMiddle, false, 0, "")

			q++
		}
	}
}

func (s Server) generateWinnersReport(c *gin.Context) {
	eventId := c.Query("eventId")

	valid := s.validateEventId(eventId)
	if !valid {
		c.JSON(http.StatusBadRequest, "EventId does not exist")
		s.Logger.Info("invalid event id")
		return
	}

	resp, err := s.GetEventWinners(stringToPrimitive(eventId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something is wrong with the server")
		s.Logger.Error(err)
		return
	}

	resp1, err := s.GetParticipantsInfoByEventID(stringToPrimitive(eventId))
	if err != nil {
		s.Logger.Error(err)
		return
	}
	winnerInfoArr := InitializeWinnersInfo(resp, resp1)

	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.AddPage()
	pdf.MoveTo(0, 0)
	pdf.SetFont("Arial", "B", 30)

	w, h := pdf.GetPageSize()

	drawBanner(pdf, w, h, bannerHt, "Event Winners")
	drawFooter(pdf, w, h, bannerHt)
	drawColumnAndInsertWinnersHeaders(pdf, xIndent, yIndent, 1)
	insertWinnerInfo(pdf, winnerInfoArr, xIndent, yIndent)
	// drawGrid(pdf, w, h)

	err = pdf.OutputFileAndClose("winners-info.pdf")
	if err != nil {
		s.Logger.Error(err.Error())
		return
	}

	fmt.Println(winnerInfoArr)

}
