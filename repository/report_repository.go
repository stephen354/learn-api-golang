package repository

import (
	"belajar-api/model"
	"database/sql"
)

type ReportRepository interface {
	GetTodaySummary() (model.SalesReport, error)
	GetSummaryByDateRange(startDate, endDate string) (model.SalesReport, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetTodaySummary() (model.SalesReport, error) {
	return r.GetSummaryByDateRange("CURRENT_DATE", "CURRENT_DATE")
}

func (r *reportRepository) GetSummaryByDateRange(startDate, endDate string) (model.SalesReport, error) {
	var report model.SalesReport

	// Total Revenue and Total Transaksi
	query := "SELECT COALESCE(SUM(total_amount), 0), COUNT(id) FROM transactions WHERE 1=1"
	var args []interface{}
	if startDate == "CURRENT_DATE" {
		query += " AND created_at::date = CURRENT_DATE"
	} else {
		query += " AND created_at::date >= $1 AND created_at::date <= $2"
		args = append(args, startDate, endDate)
	}

	err := r.db.QueryRow(query, args...).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return report, err
	}

	// Produk Terlaris
	bestProductQuery := `
		SELECT p.name, SUM(td.quantity) as total_qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE 1=1`

	if startDate == "CURRENT_DATE" {
		bestProductQuery += " AND t.created_at::date = CURRENT_DATE"
	} else {
		bestProductQuery += " AND t.created_at::date >= $1 AND t.created_at::date <= $2"
	}

	bestProductQuery += `
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1`

	err = r.db.QueryRow(bestProductQuery, args...).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)
	if err != nil && err != sql.ErrNoRows {
		return report, err
	}

	return report, nil
}
