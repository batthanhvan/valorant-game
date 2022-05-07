package reports

import "database/sql"

type Report struct {
	ReportCategory sql.NullString `gorm:"type:varchar(40)"`
	ReportDetail   sql.NullString `gorm:"type:varchar(100)"`
	ReportDate     sql.NullString `gorm:"-:migration;->"`
	RecordLink     sql.NullString `gorm:"-:migration;->"`
	UserName       sql.NullString `gorm:"type:varchar(16)"`
}
