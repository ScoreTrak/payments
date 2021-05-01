package main

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Sample struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Team      int
	Points    uint64
}

func initTables(db *gorm.DB) error {
	err := db.AutoMigrate(&Sample{})
	if err != nil {
		return err
	}
	return nil
}

func getMaxPointsPerTeam(db *gorm.DB) map[int]uint64 {
	rows, err := db.Raw("SELECT team, MAX(points) FROM samples GROUP BY team").Rows()
	if err != nil {
		panic(err) // TODO: Change this?
	}
	teamPoints := make(map[int]uint64)
	var team int
	var points uint64
	for rows.Next() {
		err = rows.Scan(&team, &points)
		if err != nil {
			panic(err) // TODO: Change this?
		}
		teamPoints[team] = points
	}
	return teamPoints
}

func updateMaxPointsPerTeam(db *gorm.DB, teamPoints map[int]uint64) {
	var samples []*Sample
	for team, points := range teamPoints {
		samples = append(samples, &Sample{Team: team, Points: points})
	}
	result := db.Create(samples)
	if result.Error != nil {
		panic(result.Error)
	}
	rowsAffected := result.RowsAffected
	log.Printf("Rows affected: %d\n", rowsAffected)
}
