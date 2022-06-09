package leaderboard

import "database/sql"

type Leaderboard struct {
	UserName      sql.NullString `gorm:"type:varchar(16)"`
	PlayerName    sql.NullString `gorm:"type:varchar(16)"`
	PlayerTagline sql.NullString `gorm:"type:bigint"`
	PlayerRank    sql.NullString `gorm:"type:varchar(13)"`
	Rating        int32          `gorm:"type:bigint"`
}
