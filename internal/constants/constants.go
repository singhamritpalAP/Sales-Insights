package constants

import "path/filepath"

var (
	CSVFilePath  = filepath.Join("..", "data", "sales_data.csv")
	DatabaseName = filepath.Join("..", "sales_database.db")
)

const (
	APIServerPort = ":8080"
	DateFormat    = "2006-01-02" // YYYY-MM-DD
	CronTime      = "@daily"
)

// query params
const (
	StartDate = "start_date"
	EndDate   = "end_date"
	Limit     = "n"
)
