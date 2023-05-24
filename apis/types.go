package apis

import (
	"github.com/sirupsen/logrus"

	"github.com/ini8labs/lsdb"
)

type Server struct {
	*logrus.Logger
	*lsdb.Client
	Addr string
}

type Date struct {
	Day   int `json:"day,omitempty"`
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}

type EventsInfo struct {
	EventUID      string `json:"event_id,omitempty"`
	EventDate     Date   `json:"event_date,omitempty"`
	EventName     string `json:"name,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	WinningNumber []int  `json:"winning_number,omitempty"`
}

type GenerateEventsInfoPDF struct {
}
