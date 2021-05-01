package report

import (
	"log"
)

type Report struct {
	Round uint
	Teams map[string]*Team
}

func (r *Report) PointsPerTeam() map[int]uint {
	points := make(map[int]uint)

	for _, team := range r.Teams {
		if !team.Pause {
			log.Println("Skipping disabled team " + team.Name)
			continue
		}

		teamNumber, err := team.Number()
		if err != nil {
			log.Println("Skipping non-numeric team " + team.Name)
			continue
		}

		points[teamNumber] = team.TotalPoints()
	}

	return points
}
