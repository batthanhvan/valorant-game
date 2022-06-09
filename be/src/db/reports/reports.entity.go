package reports

import "database/sql"

type Report struct {
	ReportCategory sql.NullString `gorm:"type:varchar(40)"`
	ReportDetail   sql.NullString `gorm:"type:varchar(100)"`
	ReportDate     sql.NullString `gorm:"-:migration;->"`
	RecordLink     sql.NullString `gorm:"-:migration;->"`
	UserName       sql.NullString `gorm:"type:varchar(16)"`
}

type PostReport struct {
	Username       string `json:"username" binding:"required"`
	ReportCategory string `json:"reportCategory" binding:"required"`
	MatchID        string `json:"matchid" binding:"required"`
	ReportDetail   string `json:"detail" binding:"required"`
}
