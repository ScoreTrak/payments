package report

import (
	"regexp"
	"strconv"
)

type Team struct {
	Name  string
	Pause bool
	Hosts map[string]*Host
}

func (t *Team) TotalPoints() uint {
	if !t.Pause {
		return 0
	}
	var total uint = 0
	for _, host := range t.Hosts {
		total += host.TotalPoints()
	}
	return total
}

// Number attempts to extract a number from the Team's Name.
func (t *Team) Number() (int, error) {
	re := regexp.MustCompile("[1-9][0-9]*")
	return strconv.Atoi(re.FindString(t.Name))
}
