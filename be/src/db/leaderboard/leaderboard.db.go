package leaderboard

import (
	"database/sql"

	"github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/lib"
	_ "github.com/go-sql-driver/mysql"
)

func PlayerCount() (*int64, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	total := db.QueryRow(CountPlayerLeaderboard)

	var r int64
	err = total.Scan(&r)
	if err != nil {
		panic(err.Error())
	}

	return &r, nil
}

func GetLeaderboardDetail(search *db.Search) ([]*pb.LeaderboardPlayer, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query(LeaderboardDetail, search.Skip, search.Limit)
	if err != nil {
		panic(err.Error())
	}

	var leaderboard Leaderboard
	rr := make([]Leaderboard, 0)
	for results.Next() {
		err = results.Scan(&leaderboard.UserName, &leaderboard.PlayerName, &leaderboard.PlayerTagline,
			&leaderboard.PlayerRank, &leaderboard.Rating)

		if err != nil {
			panic(err.Error())
		}
		rr = append(rr, leaderboard)
	}

	arr := make([]*pb.LeaderboardPlayer, 0)
	for _, v := range rr {
		arr = append(arr, ConvertMatchToProto(v))
	}
	return arr, nil
}

func ConvertMatchToProto(p Leaderboard) *pb.LeaderboardPlayer {
	ppb := &pb.LeaderboardPlayer{}

	ppb.UserName = p.UserName.String
	ppb.PlayerName = p.PlayerName.String
	ppb.PlayerTagline = p.PlayerTagline.String
	ppb.PlayerRank = p.PlayerRank.String
	ppb.Rating = p.Rating

	return ppb
}
