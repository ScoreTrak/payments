package main

import (
	"database/sql"
	"log"
	"strings"
)

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
	var values []interface{}
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
