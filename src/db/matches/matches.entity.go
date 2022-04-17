package matches

import "database/sql"

type Match struct {
	MatchID     sql.NullString `gorm:"type:varchar(12)"`
	MatchServer sql.NullString `gorm:"type:varchar(20)"`
	StartTime   sql.NullString `gorm:"type:datetime"`
	EndTime     sql.NullString `gorm:"type:datetime"`
	RecordLink  sql.NullString `gorm:"type:text"`
	MapName     sql.NullString `gorm:"type:varchar(15)"`
	ModeName    sql.NullString `gorm:"type:varchar(15)"`
	UserName    sql.NullString `gorm:"-:migration;->"`
	Agent       sql.NullString `gorm:"type:varchar(20)"`
}
