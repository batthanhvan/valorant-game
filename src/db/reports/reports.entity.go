package reports

import "database/sql"

type Report struct {
	ReportCategory sql.NullString `gorm:"type:varchar(40)"`
	ReportDetail   sql.NullString `gorm:"type:varchar(100)"`
	MatchId        sql.NullString `gorm:"type:varchar(12)"`
	UserName       sql.NullString `gorm:"type:varchar(16)"`
}
