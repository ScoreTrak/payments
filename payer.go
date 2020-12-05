package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ubnetdef/payment-processor/report"
)

const BANK_USERNAME = "username"
const BANK_PASSWORD = "password"
const BANK_BASE_URL = "https://bank.ubnetdef.org"

const BANK_AMOUNT_PER_POINT = 10

type Auth struct {
	Code   int
	Expire time.Time
	Token  string
}

func getAuth(client http.Client) (*Auth, error) {
	const loginUrl = "https://engine.ubnetdef.org/auth/login"

	form := url.Values{
		"username": {""},
		"password": {""},
	}
	resp, err := client.PostForm(loginUrl, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	auth := new(Auth)
	err = json.NewDecoder(resp.Body).Decode(auth)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func getReport(client http.Client, token string) (*report.Report, error) {
	const reportUrl = "https://engine.ubnetdef.org/api/report/"

	req, err := http.NewRequest(http.MethodGet, reportUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	report := new(report.Report)
	err = json.NewDecoder(resp.Body).Decode(report)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func initTables(db *sql.DB) {
	db.Exec(`CREATE TABLE samples (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sampled_at DATETIME DEFAULT (datetime('now')),
		team INTEGER,
		points INTEGER
	)`)
}

func getMaxPointsPerTeam(db *sql.DB) map[int]uint {
	rows, err := db.Query("SELECT team, MAX(points) FROM samples GROUP BY team")
	if err != nil {
		panic(err) // TODO: Change this?
	}

	teamPoints := make(map[int]uint)

	var team int
	var points uint

	for rows.Next() {
		err = rows.Scan(&team, &points)
		if err != nil {
			panic(err) // TODO: Change this?
		}
		teamPoints[team] = points
	}

	return teamPoints
}

func updateMaxPointsPerTeam(db *sql.DB, teamPoints map[int]uint) {
	stmtStr := "INSERT INTO samples (team, points) VALUES"
	values := []interface{}{}

	for team, points := range teamPoints {
		stmtStr += " (?, ?),"
		values = append(values, team, points)
	}
	stmtStr = strings.TrimSuffix(stmtStr, ",") // Remove extra comma.

	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(values...)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	log.Printf("Rows affected: %d\n", rowsAffected)
}

func computePointsDiff(currentPoints, previousPoints map[int]uint) map[int]uint {
	diff := make(map[int]uint)
	for team, current := range currentPoints {
		previous := previousPoints[team]
		diff[team] = current - previous
		log.Printf("Team %d: %d - %d = %d\n", team, current, previous, diff[team])
	}
	return diff
}

type Deposit struct {
	Team   int    `json:"team"`
	Amount string `json:"amount"`
}

func makeDeposit(client http.Client, teamPoints map[int]uint) {
	const depositUrl = BANK_BASE_URL + "/api/bank/accounts/deposits/"

	deposits := []Deposit{}
	for team, points := range teamPoints {
		deposit := Deposit{
			Team:   team,
			Amount: fmt.Sprintf("%d.00", points*BANK_AMOUNT_PER_POINT),
		}
		deposits = append(deposits, deposit)
	}

	data := map[string]interface{}{
		"deposits":    deposits,
		"description": "Payment for service uptime",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", depositUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(BANK_USERNAME, BANK_PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Failed to deposit")
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Missing database filename")
	}
	filename := os.Args[1]
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		panic(err) // TODO: Change this?
	}
	defer db.Close()
	initTables(db)

	client := http.Client{Timeout: 5 * time.Second}

	log.Println("Requesting authentication token")
	auth, err := getAuth(client)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Authentication token is valid until " + auth.Expire.String())

	log.Println("Requesting report")
	report, err := getReport(client, auth.Token)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Got report for round %d\n", report.Round)

	current := report.PointsPerTeam()
	previous := getMaxPointsPerTeam(db)
	diff := computePointsDiff(current, previous)
	makeDeposit(client, diff)
	updateMaxPointsPerTeam(db, current)
}
