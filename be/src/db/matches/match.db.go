package matches

import (
	"database/sql"

	"github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/lib"
	_ "github.com/go-sql-driver/mysql"
)

func MatchCount(search *db.Search) (*int64, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// total := db.QueryRow("select count(*) from playerinmatch WHERE username like ? LIMIT ?,?", search.Query, search.Skip, search.Limit)
	total := db.QueryRow("select count(*) from playerinmatch WHERE username like ?", search.Query)

	var r int64
	err = total.Scan(&r)
	if err != nil {
		panic(err.Error())
	}

	return &r, nil
}

func MatchesByUsername(search *db.Search) ([]*pb.Match, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query(MatchesDetailQuery, search.Query, search.Skip, search.Limit)
	if err != nil {
		panic(err.Error())
	}

	var r Match
	rr := make([]Match, 0)
	for results.Next() {
		err = results.Scan(&r.MatchID, &r.MatchServer, &r.MapName, &r.ModeName, &r.StartTime, &r.EndTime, &r.RecordLink)

		if err != nil {
			panic(err.Error())
		}
		rr = append(rr, r)
	}

	arr := make([]*pb.Match, 0)
	for _, v := range rr {
		arr = append(arr, ConvertMatchToProto(v))
	}
	return arr, nil

}

func ConvertMatchToProto(p Match) *pb.Match {
	ppb := &pb.Match{}

	ppb.MatchID = p.MatchID.String
	ppb.MatchServer = p.MatchServer.String
	ppb.MapName = p.MapName.String
	ppb.ModeName = p.ModeName.String
	ppb.StartTime = p.StartTime.String
	ppb.EndTime = p.EndTime.String
	ppb.RecordLink = p.RecordLink.String

	return ppb
}
