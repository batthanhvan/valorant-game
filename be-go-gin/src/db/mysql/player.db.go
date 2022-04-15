package mysql

import (
	"database/sql"

	"github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/players"
	_ "github.com/go-sql-driver/mysql"
)

func TotalPlayer(search *db.Search) (*int64, error) {

	db, err := sql.Open("mysql", db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	search.Query = "%" + search.Query + "%"
	total := db.QueryRow("SELECT count(sub.username) FROM (SELECT players.username,reports.reportCategory,reports.reportDetail FROM players LEFT JOIN reports ON players.username = reports.username WHERE players.username like ? LIMIT ?,?) as sub", search.Query, search.Skip, search.Limit)

	var r int64
	err = total.Scan(&r)
	if err != nil {
		panic(err.Error())
	}

	return &r, nil
}

func ListPlayers(search *db.Search) ([]*pb.Player, error) {

	db, err := sql.Open("mysql", db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	search.Query = "%" + search.Query + "%"
	results, err := db.Query("SELECT players.username,reports.reportCategory,reports.reportDetail FROM players LEFT JOIN reports ON players.username = reports.username WHERE players.username like ? LIMIT ?,?", search.Query, search.Skip, search.Limit)

	if err != nil {
		panic(err.Error())
	}

	var r players.Player
	rr := make([]players.Player, 0)
	for results.Next() {

		err = results.Scan(&r.UserName, &r.ReportCategory, &r.ReportDetail)

		if err != nil {
			panic(err.Error())
		}
		rr = append(rr, r)
	}

	arr := make([]*pb.Player, 0)
	for _, v := range rr {
		arr = append(arr, ConvertPlayerToProto(v))
	}
	return arr, nil

}

func ConvertPlayerToProto(p players.Player) *pb.Player {
	ppb := &pb.Player{}

	ppb.UserName = p.UserName
	ppb.ReportCategory = p.ReportCategory
	ppb.ReportDetail = p.ReportDetail

	return ppb
}
