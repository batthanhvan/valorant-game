package db

const ConStr = "root:1234567@/valorant"

type Search struct {
	Skip      int
	Limit     int
	OrderBy   string
	OrderType string
	Query     string
	// Fields    []string
}
