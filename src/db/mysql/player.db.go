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

	search.Query = "%" + search.Query + "%"
	total := db.QueryRow("SELECT count(username) FROM players WHERE username like ? LIMIT ?,?", search.Query, search.Skip, search.Limit)

	var r int64
	err = total.Scan(&r)
	if err != nil {
		panic(err.Error())
	}

	return &r, nil
}

func SearchPlayer(search *db.Search) ([]*pb.Player, error) {

	db, err := sql.Open("mysql", db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	search.Query = "%" + search.Query + "%"
	results, err := db.Query(players.QueryString, search.Query, search.Skip, search.Limit)

	if err != nil {
		panic(err.Error())
	}

	var player players.Player
	rr := make([]players.Player, 0)
	for results.Next() {

		err = results.Scan(&player.PlayerName, &player.PlayerTagline, &player.PlayerRank, &player.PlayerStatus,
			&player.Wins, &player.Kills, &player.Assists, &player.KillsPerRound, &player.FirstBloods, &player.Aces,
			&player.Clutches, &player.MostKills)

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

	ppb.PlayerName = p.PlayerName.String
	ppb.PlayerTagline = p.PlayerTagline
	ppb.PlayerRank = p.PlayerRank.String
	ppb.PlayerStatus = p.PlayerStatus.String
	ppb.Wins = p.Wins
	ppb.Kills = p.Kills
	ppb.Assists = p.Assists
	ppb.KillsPerRound = p.KillsPerRound
	ppb.FirstBloods = p.FirstBloods
	ppb.Aces = p.Aces
	ppb.Clutches = p.Clutches
	ppb.MostKills = p.MostKills

	return ppb
}
