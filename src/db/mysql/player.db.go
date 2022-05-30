package mysql

import (
	"database/sql"
	"strings"

	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/players"
	"github.com/batthanhvan/src/lib"
	_ "github.com/go-sql-driver/mysql"
)

// func PlayerCount(search *db.Search) (*int64, error) {

// 	db, err := sql.Open("mysql", db.ConStr)
// 	lib.CheckError(err)
// 	defer db.Close()

// 	search.Query = "%" + search.Query + "%"
// 	total := db.QueryRow("SELECT count(username) FROM players WHERE username like ? LIMIT ?,?", search.Query, search.Skip, search.Limit)

// 	var r int64
// 	err = total.Scan(&r)
// 	lib.CheckError(err)

// 	return &r, nil
// }

func GetPlayer(search *db.Search) (*pb.Player, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	lib.CheckError(err)
	defer db.Close()

	results, err := db.Query(players.GetPlayerDetailQuery, search.Query)

	lib.CheckError(err)

	var player players.Player
	//	rr := make([]players.Player, 0)
	for results.Next() {

		err = results.Scan(&player.UserName, &player.PlayerName, &player.PlayerTagline, &player.PlayerRank,
			&player.PlayerStatus, &player.Wins, &player.Kills, &player.Assists, &player.KillsPerRound,
			&player.FirstBloods, &player.Aces, &player.Clutches, &player.MostKills)

		lib.CheckError(err)
		//rr = append(rr, player)
	}

	// arr := make([]*pb.Player, 0)
	// for _, player := range rr {
	// 	arr = append(arr, ConvertPlayerToProto(player))
	// }
	return ConvertPlayerToProto(player), nil
}

func ModifyPlayer(username string, playername string, tagline string) (*pb.Player, string, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	lib.CheckError(err)
	defer db.Close()

	checkUsernameExists, err := db.Query("SELECT 1 FROM players WHERE username=?", username)
	lib.CheckError(err)

	resultStatus := ""
	if checkUsernameExists.Next() {
		resultStatus = modifyPlayerName(db, username, playername, resultStatus)
		resultStatus = modifyTagline(db, username, tagline, resultStatus)
	} else {
		resultStatus = "Username not exists. "
	}

	result, err := db.Query(players.GetPlayerDetailQuery, username)
	lib.CheckError(err)

	var player players.Player
	for result.Next() {
		err = result.Scan(&player.UserName, &player.PlayerName, &player.PlayerTagline, &player.PlayerRank,
			&player.PlayerStatus, &player.Wins, &player.Kills, &player.Assists, &player.KillsPerRound,
			&player.FirstBloods, &player.Aces, &player.Clutches, &player.MostKills)
		lib.CheckError(err)
	}

	return ConvertPlayerToProto(player), resultStatus, nil
}

func modifyPlayerName(db *sql.DB, username string, playername string, resultStatus string) string {

	if playername == "" {
		resultStatus += "Empty player name. "
	} else if strings.Contains(playername, " ") {
		resultStatus += "Player name can't contain space character. "
	} else {
		checkPlayernameInUse, err := db.Query("SELECT 1 FROM players WHERE username=? AND playerName=?", username, playername)
		lib.CheckError(err)

		if checkPlayernameInUse.Next() {
			resultStatus += "Player name is in use. "
		} else {
			_, err := db.Query(players.ModifyPlayerNameQuery, playername, username)
			lib.CheckError(err)
			resultStatus += "Player name is changed successfully. "
		}
	}
	return resultStatus
}

func modifyTagline(db *sql.DB, username string, tagline string, resultStatus string) string {
	if tagline == "" {
		resultStatus += "Empty tagline. "
	} else if strings.Contains(tagline, " ") {
		resultStatus += "Tagline can't contain space character. "
	} else {
		checkTaglineInUse, err := db.Query("SELECT 1 FROM players WHERE username=? AND playerTagline=?", username, tagline)
		lib.CheckError(err)

		if checkTaglineInUse.Next() {
			resultStatus += "Tagline is in use. "
		} else {
			_, err = db.Query(players.ModifyTaglineQuery, tagline, username)
			lib.CheckError(err)
			resultStatus += "Tagline is changed successfully. "
		}
	}
	return resultStatus
}

func ConvertPlayerToProto(p players.Player) *pb.Player {
	ppb := &pb.Player{}

	ppb.Username = p.UserName.String
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
