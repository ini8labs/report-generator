package apis

import (
	"strconv"
	"strings"

	apis "github.com/ini8labs/admin-service/src/apis"
	"github.com/jung-kurt/gofpdf"
)

const (
	bannerHt = 43
	xIndent  = 5
	yIndent  = 47
)

func convertDateToString(date apis.Date) string {
	dateArray := []int{date.Day, date.Month, date.Year}
	strArray := make([]string, len(dateArray))
	for i, num := range dateArray {
		strArray[i] = strconv.Itoa(num)
	}
	str := strings.Join(strArray, "/")
	return str
}

func convertWinNumbersToString(numbers []int) string {
	strArray := make([]string, len(numbers))
	for i, num := range numbers {
		strArray[i] = strconv.Itoa(num)
	}
	str := strings.Join(strArray, ",")
	return str
}
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
