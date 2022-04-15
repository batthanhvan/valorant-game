package players

type Player struct {
	UserName       string `gorm:"type:varchar(32)"`
	PlayerName     string `gorm:"type:varchar(32)"`
	PlayerTagline  int64  `gorm:"type:bigint"`
	PlayerRank     string `gorm:"type:varchar(32)"`
	PlayerStatus   string `gorm:"type:varchar(32)"`
	ReportCategory string `gorm:"-:migration;->"`
	ReportDetail   string `gorm:"-:migration;->"`
}
