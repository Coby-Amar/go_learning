package handlers

import (
	"log/slog"
	"net/http"
)

func (conf *ApiConfig) HandleGetReports(w http.ResponseWriter, r *http.Request) {
	slog.Info("HandleGetReports")
	sessionParams := conf.getSessionParams(r)
	if sessionParams == nil {
		respondWithInternalServerError(w)
		return
	}
	reports, err := conf.DB.GetAllUserReports(r.Context(), sessionParams.UserID)
	if err != nil {
		slog.Error("DB error GetAllReports", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	respondWithJSON(w, http.StatusOK, reports)
}

func (conf *ApiConfig) HandleCreateReport(w http.ResponseWriter, r *http.Request, params UserCreateReportWithEntries) {
	slog.Info("HandleCreateReport")
	sessionParams := conf.getSessionParams(r)
	if sessionParams == nil {
		respondWithInternalServerError(w)
		return
	}
	params.Report.UserID = sessionParams.UserID
	dbCreatedReport, err := conf.DB.CreateReport(r.Context(), params.Report)
	if err != nil {
		slog.Error("DB CreateReport", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	for _, entry := range params.Entries {
		entry.ReportID = dbCreatedReport.ID
	}
	_, err = conf.DB.CreateReportEntries(r.Context(), params.Entries)
	if err != nil {
		slog.Error("DB CreateReportEntries", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	respondWithJSON(w, http.StatusCreated, dbCreatedReport)
}
