package report

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"regexp"
	"strconv"
)

func TotalTeamPoints(t *report.SimpleTeam) uint64 {
	if t.Pause {
		return 0
	}
	var total uint64 = 0
	for _, host := range t.Hosts {
		total += TotalHostPoints(host)
	}
	return total
}

// TeamNumber attempts to extract a number from the Team's Name.
func TeamNumber(t *report.SimpleTeam) (int, error) {
	re := regexp.MustCompile("[1-9][0-9]*")
	return strconv.Atoi(re.FindString(t.Name))
}
