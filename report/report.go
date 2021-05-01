package report

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"log"
)

func PointsPerTeam(r *report.SimpleReport) map[int]uint64 {
	points := make(map[int]uint64)
	for _, team := range r.Teams {
		if team.Pause {
			log.Println("Skipping disabled team " + team.Name)
			continue
		}
		teamNumber, err := TeamNumber(team)
		if err != nil {
			log.Println("Skipping non-numeric team " + team.Name)
			continue
		}

		points[teamNumber] = TotalTeamPoints(team)
	}
	return points
}
