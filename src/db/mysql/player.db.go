package mysql

import (
	"database/sql"

	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/players"
	_ "github.com/go-sql-driver/mysql"
)

func PlayerCount(search *db.Search) (*int64, error) {

	db, err := sql.Open("mysql", db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	total := db.QueryRow("select count(*) from players WHERE players.username like ? LIMIT ?,?", "%"+search.Query+"%", search.Skip, search.Limit)

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

	results, err := db.Query("SELECT players.username,reports.reportCategory,reports.reportDetail FROM players LEFT JOIN reports ON players.username = reports.username WHERE players.username like ? LIMIT ?,?", "%"+search.Query+"%", search.Skip, search.Limit)

	if err != nil {
		panic(err.Error())
	}

	var player players.Player
	rr := make([]players.Player, 0)
	for results.Next() {

		err = results.Scan(&player.UserName, &player.ReportCategory, &player.ReportDetail)

		if err != nil {
			panic(err.Error())
		}
		rr = append(rr, player)
	}

	arr := make([]*pb.Player, 0)
	for _, player := range rr {
		arr = append(arr, ConvertPlayerToProto(player))
	}
	return arr, nil

}

func ConvertPlayerToProto(p players.Player) *pb.Player {
	ppb := &pb.Player{}

	ppb.UserName = p.UserName.String
	ppb.ReportCategory = p.ReportCategory.String
	ppb.ReportDetail = p.ReportDetail.String

	return ppb
}
