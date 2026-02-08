package handler

import (
	"belajar-api/service"
	"encoding/json"
	"net/http"
)

type ReportHandler struct {
	service service.ReportService
}

func NewReportHandler(service service.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	var report interface{}
	var err error

	if startDate != "" && endDate != "" {
		report, err = h.service.GetSummaryByDateRange(startDate, endDate)
	} else {
		report, err = h.service.GetTodaySummary()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
