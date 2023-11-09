package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/main/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

// ---------------------------------------------------------------------------------------------------
// HandleGetReports
// ---------------------------------------------------------------------------------------------------
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

// ---------------------------------------------------------------------------------------------------
// HandleGetReportEntries
// ---------------------------------------------------------------------------------------------------
func HandleGetReportEntries(cwrar *utils.ConfigWithRequestAndResponse) {
	slog.Info("HandleGetReportEntries")
	reportId, err := utils.GetIdFromURLParam(cwrar.R, utils.REPORT_ID)
	if err != nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	reportEntries, err := cwrar.Config.Queries.GetReportEntries(cwrar.R.Context(), reportId)
	if err != nil {
		slog.Error("GetAllReports", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, reportEntries)
}

// ---------------------------------------------------------------------------------------------------
// createReportEntries
// ---------------------------------------------------------------------------------------------------
func createReportEntries(queries *database.Queries, context context.Context, entries []database.CreateReportEntriesParams, reportId pgtype.UUID) error {
	for index := range entries {
		entries[index].ReportID = reportId
	}
	_, err := queries.CreateReportEntries(context, entries)
	if err != nil {
		return err
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------
// HandleCreateReport
// ---------------------------------------------------------------------------------------------------
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
	err = createReportEntries(queries, context, params.Entries, dbCreatedReport.ID)
	if err != nil {
		slog.Error("CreateReportEntries", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	tx.Commit(context)
	utils.RespondWithJSON(cwrar.W, http.StatusCreated, dbCreatedReport)
}

// ---------------------------------------------------------------------------------------------------
// HandleUpdateReport
// ---------------------------------------------------------------------------------------------------
func HandleUpdateReport(cwrar *utils.ConfigWithRequestAndResponse, params utils.UserUpdateReportWithEntries) {
	params.Report.AmoutOfEntries = int16(len((params.ExistingEntries)) + len(params.EntriesToCreate))

	context := cwrar.R.Context()
	tx, err := cwrar.Config.Connection.Begin(context)
	if err != nil {
		slog.Info("Failed to Begin Transaction", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	defer tx.Rollback(context)

	queries := cwrar.Config.Queries.WithTx(tx)
	dbUpdatedReport, updateErr := queries.UpdateReport(context, params.Report)
	if updateErr != nil {

		slog.Error("UpdateReport", utils.ERROR, updateErr)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	for _, entry := range params.ExistingEntries {
		_, err = queries.UpdateReportEntry(context, entry)
		if err != nil {
			slog.Error("UpdateReportEntry", utils.ERROR, err)
			utils.RespondWithInternalServerError(cwrar.W)
			return
		}
	}
	err = createReportEntries(queries, context, params.EntriesToCreate, dbUpdatedReport.ID)
	if err != nil {
		slog.Error("CreateReportEntries", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	tx.Commit(context)
	utils.RespondWithJSON(cwrar.W, http.StatusAccepted, dbUpdatedReport)
}

// ---------------------------------------------------------------------------------------------------
// HandleDeleteReport
// ---------------------------------------------------------------------------------------------------
func HandleDeleteReport(cwrar *utils.ConfigWithRequestAndResponse) {
	slog.Error("HandleDeleteReport")
	reportId, err := utils.GetIdFromURLParam(cwrar.R, utils.REPORT_ID)
	if err != nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	delteErr := cwrar.Config.Queries.DeleteReport(cwrar.R.Context(), reportId)
	if delteErr != nil {
		slog.Error("DeleteReport", utils.ERROR, delteErr)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithMessage(cwrar.W, http.StatusOK, "deleted")
}
