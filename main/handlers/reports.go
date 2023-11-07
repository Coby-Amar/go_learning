package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/main/utils"
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

func HandleGetReportEntries(cwrar *utils.ConfigWithRequestAndResponse) {
	slog.Info("HandleGetReportEntries")
	reportId, err := utils.GetIdFromURLParam(cwrar.R, utils.REPORT_ID)
	if err != nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	reportEtries, err := cwrar.Config.Queries.GetReportEntries(cwrar.R.Context(), reportId)
	if err != nil {
		slog.Error("GetAllReports", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, reportEtries)
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

func HandleUpdateReport(cwrar *utils.ConfigWithRequestAndResponse, params utils.UserUpdateReportWithEntries) {
	slog.Info("HandleUpdateReport")
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
	dbUpdatedReport, updateErr := queries.UpdateReport(context, params.Report)
	if updateErr != nil {
		slog.Error("UpdateReport", utils.ERROR, updateErr)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	entriesToCreate := []database.CreateReportEntriesParams{}
	for _, entry := range params.Entries {
		if entry.ID.Valid {
			_, err = queries.UpdateReportEntry(context, database.UpdateReportEntryParams{
				ID:            entry.ID,
				Amount:        entry.Amount,
				Carbohydrates: entry.Carbohydrates,
				Proteins:      entry.Proteins,
				Fats:          entry.Fats,
			})
			if err != nil {
				slog.Error("UpdateReportEntry", utils.ERROR, err)
				utils.RespondWithInternalServerError(cwrar.W)
				return
			}
		} else {
			entriesToCreate = append(entriesToCreate, database.CreateReportEntriesParams{
				ProductID:     entry.ProductID,
				ReportID:      dbUpdatedReport.ID,
				Amount:        entry.Amount,
				Carbohydrates: entry.Carbohydrates,
				Proteins:      entry.Proteins,
				Fats:          entry.Fats,
			})

		}
	}
	_, err = queries.CreateReportEntries(context, entriesToCreate)
	if err != nil {
		slog.Error("CreateReportEntries", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	tx.Commit(context)
	utils.RespondWithJSON(cwrar.W, http.StatusAccepted, dbUpdatedReport)
}

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
