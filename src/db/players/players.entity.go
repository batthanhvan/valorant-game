package players

import "database/sql"

type Player struct {
	UserName       sql.NullString `gorm:"type:varchar(16)"`
	PlayerName     sql.NullString `gorm:"type:varchar(16)"`
	PlayerTagline  int64          `gorm:"type:bigint"`
	PlayerRank     sql.NullString `gorm:"type:varchar(13)"`
	PlayerStatus   sql.NullString `gorm:"type:varchar(18)"`
	ReportCategory sql.NullString `gorm:"-:migration;->"`
	ReportDetail   sql.NullString `gorm:"-:migration;->"`
}
