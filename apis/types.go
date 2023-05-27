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
