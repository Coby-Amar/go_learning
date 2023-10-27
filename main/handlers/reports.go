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
	reports, err := cwrar.Config.Queries.GetAllUserReports(cwrar.R.Context(), cwrar.Sparams.UserID)
	if err != nil {
		slog.Error("GetAllReports", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, reports)
}

func HandleCreateReport(cwrar *utils.ConfigWithRequestAndResponse, params utils.UserCreateReportWithEntries) {
	slog.Info("HandleCreateReport")
	params.Report.UserID = cwrar.Sparams.UserID
	params.Report.AmoutOfEntries = int16(len((params.Entries)))

	context := cwrar.R.Context()
	tx, err := cwrar.Config.Connection.Begin(context)
	if err != nil {
		slog.Info("Failed to Begin Transaction", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	defer tx.Rollback(context)

	queries := cwrar.Config.Queries.WithTx(tx)
	dbCreatedReport, err := queries.CreateReport(context, params.Report)
	if err != nil {
		slog.Error("CreateReport", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	for index := range params.Entries {
		params.Entries[index].ReportID = dbCreatedReport.ID
	}
	_, err = queries.CreateReportEntries(context, params.Entries)
	if err != nil {
		slog.Error("CreateReportEntries", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	tx.Commit(context)
	utils.RespondWithJSON(cwrar.W, http.StatusCreated, dbCreatedReport)
}

func HandleDeleteReport(cwrar *utils.ConfigWithRequestAndResponse) {
	slog.Error("HandleDeleteReport")
	reportId := utils.GetIdFromURLParam(cwrar.R, utils.REPORT_ID)
	if reportId == uuid.Nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	delteErr := cwrar.Config.Queries.DeleteReport(cwrar.R.Context(), pgtype.UUID{
		Bytes: reportId,
		Valid: true,
	})
	if delteErr != nil {
		slog.Error("DeleteReport", utils.ERROR, delteErr)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithMessage(cwrar.W, http.StatusOK, "deleted")
}
