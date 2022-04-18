package main

import (
	"bufio"
	"database/sql"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thanhpk/randstr"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	db := initDB()

	usernames := populatePlayers(db)
	matchIDs := populateMatches(db, usernames)
	populateReports(db, usernames, matchIDs)

	db.Close()
}

func initDB() *sql.DB {
	file, err := os.Open("src\\db\\mysql\\userpassword.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var username, password, DBname string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	username = scanner.Text()
	scanner.Scan()
	password = scanner.Text()
	scanner.Scan()
	DBname = scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var DBsource = username + ":" + password + "@/" + DBname
	db, err := sql.Open("mysql", DBsource)
	checkErr(err)
	return db
}

func populatePlayers(db *sql.DB) []string {
	var usernames []string
	statuses := []string{"Active", "Permanently banned", "Temporarily banned"}
	ranks := []string{"Unidentified", "Iron ", "Bronze ", "Silver ", "Gold ", "Platinum ", "Diamond ", "Immortal", "Radiant"}
	playersCount := 50

	stmt, err := db.Prepare("INSERT players SET username=?, playerName=?, playerTagline=?, playerRank=?, playerStatus=?")
	checkErr(err)

	for i := 0; i < playersCount; i++ {
		randUsername := randstr.String(rand.Intn(16) + 1)
		usernames = append(usernames, randUsername)
		randPlayerName := randstr.String(rand.Intn(16) + 1)
		randRank := ranks[rand.Intn(len(ranks))]
		if randRank != "Immortal" && randRank != "Radiant" && randRank != "Unidentified" {
			randRank += strconv.Itoa((rand.Intn(3) + 1))
		}
		randStatus := statuses[rand.Intn(len(statuses))]
		_, err := stmt.Exec(randUsername, randPlayerName, rand.Intn(99899)+100, randRank, randStatus)
		checkErr(err)
	}
	return usernames
}

func populateMatches(db *sql.DB, usernames []string) []string {
	var matchIDs []string
	servers := []string{"Singapore ", "Hong Kong ", "Tokyo", "Sydney ", "Mumbai"}
	maps := []string{"Bind", "Icebox", "Split", "Ascent", "Breeze", "Haven", "Fracture"}
	modes := []string{"Competitive", "Unrated", "Spike Rush", "Deathmatch", "Snowball Fight", "Replication", "Custom"}
	agents := []string{"Astra", "Breach", "Brimstone", "Chamber", "Cypher", "Jett", "Kay/O", "Killjoy", "Neon", "Omen", "Phoenix", "Raze", "Reyna", "Sage", "Skye", "Sova", "Viper", "Yoru"}
	matchesCount := 50

	stmt, err := db.Prepare("INSERT matches SET matchID=?, matchServer=?, startTime=?, endTime=?, recordLink=?, mapName=?, modeName=?, username=?, agent=?")
	checkErr(err)

	for i := 0; i < matchesCount; i++ {
		randMatchID := randstr.String(13)
		matchIDs = append(matchIDs, randMatchID)
		randServer := servers[rand.Intn(len(servers))]
		if randServer != "Mumbai" {
			randServer += strconv.Itoa((rand.Intn(2) + 1))
		}
		randTime := randate()
		randStartTime := randTime.Format("2006-01-02 15:04:05")
		randEndTime := randTime.Add(time.Minute*time.Duration(rand.Intn(41)+30) + time.Second*time.Duration(rand.Intn(60))).Format("2006-01-02 15:04:05")
		randLink := randstr.String(rand.Intn(70) + 1)
		randMap := maps[rand.Intn(len(maps))]
		randMode := modes[rand.Intn(len(modes))]
		randUser := usernames[rand.Intn(len(usernames))]
		randAgent := agents[rand.Intn(len(agents))]
		_, err := stmt.Exec(randMatchID, randServer, randStartTime, randEndTime, randLink, randMap, randMode, randUser, randAgent)
		checkErr(err)
	}
	return matchIDs
}

func populateReports(db *sql.DB, usernames []string, matchIDs []string) {
	reportCatagories := []string{"AFK", "Assisting Enemy", "Cheating", "Communication Abuse - Text", "Communication Abuse - Voice", "Negative Attitude", "Offensive Name", "Threats"}
	reportsCount := 50

	stmt, err := db.Prepare("INSERT reports SET reportCategory=?, reportDetail=?, matchID=?, username=?")
	checkErr(err)

	for i := 0; i < reportsCount; i++ {
		randCatagory := reportCatagories[rand.Intn(len(reportCatagories))]
		randDetail := randstr.String(rand.Intn(100) + 1)
		randMatchID := matchIDs[rand.Intn(len(matchIDs))]
		randUser := usernames[rand.Intn(len(usernames))]
		_, err := stmt.Exec(randCatagory, randDetail, randMatchID, randUser)
		checkErr(err)
	}
}

func randate() time.Time {
	min := time.Date(2020, 2, 2, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// func main() {

// }
