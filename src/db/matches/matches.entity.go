package matches

import "database/sql"

type Match struct {
	MatchID     sql.NullString `gorm:"type:varchar(20)"`
	MatchServer sql.NullString `gorm:"type:varchar(20)"`
	MapName     sql.NullString `gorm:"type:varchar(15)"`
	ModeName    sql.NullString `gorm:"type:varchar(15)"`
	StartTime   sql.NullString `gorm:"type:datetime"`
	EndTime     sql.NullString `gorm:"type:datetime"`
	RecordLink  sql.NullString `gorm:"type:text"`
}
