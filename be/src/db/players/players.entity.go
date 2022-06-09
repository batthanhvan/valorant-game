package players

import "database/sql"

type Player struct {
	UserName      sql.NullString `gorm:"type:varchar(16)"`
	PlayerName    sql.NullString `gorm:"type:varchar(16)"`
	PlayerTagline sql.NullString `gorm:"type:varchar(5)"`
	PlayerRank    sql.NullString `gorm:"type:varchar(13)"`
	PlayerStatus  sql.NullString `gorm:"type:varchar(18)"`
	Wins          int32          `gorm:"type:bigint"`
	Kills         int32          `gorm:"type:bigint"`
	Assists       int32          `gorm:"type:bigint"`
	KillsPerRound float32        `gorm:"type:float"`
	FirstBloods   int32          `gorm:"type:bigint"`
	Aces          int32          `gorm:"type:bigint"`
	Clutches      int32          `gorm:"type:bigint"`
	MostKills     int32          `gorm:"type:bigint"`
}
