package handlers

import (
	"time"

	"github.com/coby-amar/go_learning/database"
)

type userGetReport struct {
	Date    time.Time                      `json:"date"`
	Entries []database.GetReportEntriesRow `json:"entries"`
}

type userCreateReport struct {
	Date    time.Time              `json:"date"`
	Entries []database.ReportEntry `json:"entries"`
}
