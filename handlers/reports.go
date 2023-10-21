package handlers

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"sync"

	"github.com/coby-amar/go_learning/database"
)

func getReportEntries(conf *ApiConfig, ctx context.Context, dbReport database.Report, ch chan userGetReport, wg *sync.WaitGroup) {
	entries, err := conf.DB.GetReportEntries(ctx, dbReport.ID)
	if err != nil {
		log.Fatal("DB error on GetReportEntries request:", err)
		return
	}
	ch <- userGetReport{
		userReport: userReport{Date: dbReport.Date},
		Entries:    entries,
	}
	wg.Done()
}

func (conf *ApiConfig) HandleGetReports(w http.ResponseWriter, r *http.Request) {
	reports, err := conf.DB.GetAllReports(r.Context())
	if err != nil {
		slog.Error("DB error on GetAllProducts", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	reportsLength := len(reports)
	if reportsLength == 0 {
		respondWithJSON(w, http.StatusOK, reports)
		return
	}
	ch := make(chan userGetReport, reportsLength)
	wg := &sync.WaitGroup{}
	wg.Add(reportsLength)
	for _, report := range reports {
		go getReportEntries(conf, r.Context(), report, ch, wg)
	}
	wg.Wait()
	close(ch)
	froundReports := []userGetReport{}
	for result := range ch {
		froundReports = append(froundReports, result)
	}
	respondWithJSON(w, http.StatusOK, froundReports)
}

func (conf *ApiConfig) HandleCreateReport(w http.ResponseWriter, r *http.Request, params UserRequestReport) {
	dbCreatedReport, err := conf.DB.CreateReport(r.Context(), database.CreateReportParams{
		Date:           params.Date,
		AmoutOfEntries: int16(len(params.Entries)),
	})
	if err != nil {
		slog.Error("DB error on CreateReport", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	reportForUser := userCreatedReport{
		userReport: userReport{Date: dbCreatedReport.Date},
	}
	for _, entry := range params.Entries {
		entry.ReportID = dbCreatedReport.ID
		createdEntry, err := conf.DB.CreateReportEntry(r.Context(), entry)
		if err != nil {
			slog.Error("Failed to create entry", "entry", entry, ERROR, err)
			continue
		}
		reportForUser.Entries = append(reportForUser.Entries, createdEntry)
	}
	respondWithJSON(w, http.StatusCreated, reportForUser)
}
