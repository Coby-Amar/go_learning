package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
)

func HandleGetReports(cwrar *utils.ConfigWithRequestAndResponse) {
	slog.Info("HandleGetReports")
	sessionParams := utils.GetSessionParams(cwrar)
	if sessionParams == nil {
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	reports, err := cwrar.Config.DB.GetAllUserReports(cwrar.R.Context(), sessionParams.UserID)
	if err != nil {
		slog.Error("DB error GetAllReports", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, reports)
}

func HandleCreateReport(cwrar *utils.ConfigWithRequestAndResponse, params *UserCreateReportWithEntries) {
	slog.Info("HandleCreateReport")
	sessionParams := utils.GetSessionParams(cwrar)
	if sessionParams == nil {
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	params.Report.UserID = sessionParams.UserID
	dbCreatedReport, err := cwrar.Config.DB.CreateReport(cwrar.R.Context(), params.Report)
	if err != nil {
		slog.Error("DB CreateReport", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	for _, entry := range params.Entries {
		entry.ReportID = dbCreatedReport.ID
	}
	_, err = cwrar.Config.DB.CreateReportEntries(cwrar.R.Context(), params.Entries)
	if err != nil {
		slog.Error("DB CreateReportEntries", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusCreated, dbCreatedReport)
}
