package reports

type Player struct {
	ReportCategory string `gorm:"type:varchar(32)"`
	ReportDetail   string `gorm:"type:varchar(64)"`
	MatchId        string `gorm:"type:varchar(32)"`
	UserName       string `gorm:"type:varchar(32)"`
}
