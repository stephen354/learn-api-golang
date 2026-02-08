package service

import (
	"belajar-api/model"
	"belajar-api/repository"
)

type ReportService interface {
	GetTodaySummary() (model.SalesReport, error)
	GetSummaryByDateRange(startDate, endDate string) (model.SalesReport, error)
}

type reportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) GetTodaySummary() (model.SalesReport, error) {
	return s.repo.GetTodaySummary()
}

func (s *reportService) GetSummaryByDateRange(startDate, endDate string) (model.SalesReport, error) {
	return s.repo.GetSummaryByDateRange(startDate, endDate)
}
