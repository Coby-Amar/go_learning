package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func HandleCreateReport(cwrar *utils.ConfigWithRequestAndResponse, params utils.UserCreateReportWithEntries) {
	sessionParams := utils.GetSessionParams(cwrar)
	if sessionParams == nil {
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	params.Report.UserID = sessionParams.UserID
	params.Report.AmoutOfEntries = int16(len((params.Entries)))
	dbCreatedReport, err := cwrar.Config.DB.CreateReport(cwrar.R.Context(), params.Report)
	if err != nil {
		slog.Error("DB CreateReport", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	for index := range params.Entries {
		params.Entries[index].ReportID = dbCreatedReport.ID
	}
	_, err = cwrar.Config.DB.CreateReportEntries(cwrar.R.Context(), params.Entries)
	if err != nil {
		slog.Error("DB CreateReportEntries", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusCreated, dbCreatedReport)
}

func HandleDeleteReport(cwrar *utils.ConfigWithRequestAndResponse) {
	reportId := utils.GetIdFromURLParam(cwrar.R, utils.REPORT_ID)
	if reportId == uuid.Nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	delteErr := cwrar.Config.DB.DeleteReport(cwrar.R.Context(), pgtype.UUID{
		Bytes: reportId,
		Valid: true,
	})
	if delteErr != nil {
		slog.Error("HandleDeleteReport DeleteReport", utils.ERROR, delteErr)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithMessage(cwrar.W, http.StatusOK, "deleted")
}
