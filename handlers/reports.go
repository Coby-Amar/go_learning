package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/coby-amar/go_learning/database"
)

func getReportEntries(conf *ApiConfig, ctx context.Context, dbReport database.Report, ch chan userGetReport, wg *sync.WaitGroup) {
	entries, err := conf.DB.GetReportEntries(ctx, dbReport.ID)
	if err != nil {
		log.Fatal("DB error on GetReportEntries request:", err)
		return
	}
	ch <- userGetReport{
		Date:    dbReport.Date,
		Entries: entries,
	}
	wg.Done()
}

func (conf *ApiConfig) HandleGetReports(w http.ResponseWriter, r *http.Request) {
	reports, err := conf.DB.GetAllReports(r.Context())
	if err != nil {
		log.Fatal("DB error on GetAllProducts request:", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve all products")
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

func (conf *ApiConfig) HandleCreateReport(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type userReport struct {
		Date    time.Time                          `json:"date"`
		Entries []database.CreateReportEntryParams `json:"entries"`
	}
	params := userReport{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("CreateReportWithEntries error: ", err)
		respondWithError(w, 400, "Given parameters are invalid or missing")
		return
	}
	log.Println(params.Date)
	createdReport, err := conf.DB.CreateReport(r.Context(), params.Date)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating product")
		log.Println(err)
		return
	}
	reportForUser := userCreateReport{
		Date: createdReport.Date,
	}
	for _, entry := range params.Entries {
		entry.ReportID = createdReport.ID
		createdEntry, err := conf.DB.CreateReportEntry(r.Context(), entry)
		if err != nil {
			log.Println(err)
			continue
		}
		reportForUser.Entries = append(reportForUser.Entries, createdEntry)
	}
	respondWithJSON(w, http.StatusCreated, reportForUser)
}
