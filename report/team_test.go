package report

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/gofrs/uuid"
	"testing"
)

func TestTeamEnabled(t *testing.T) {
	team := report.SimpleTeam{
		Pause: false,
		Hosts: map[uuid.UUID]*report.SimpleHost{
			uuid.Must(uuid.NewV4()): {
				Pause: false,
				Services: map[uuid.UUID]*report.SimpleService{
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 1},
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 20},
				},
			},
			uuid.Must(uuid.NewV4()): {
				Pause: true,
				Services: map[uuid.UUID]*report.SimpleService{
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 300},
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 4000},
				},
			},
			uuid.Must(uuid.NewV4()): {
				Pause: false,
				Services: map[uuid.UUID]*report.SimpleService{
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 50000},
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 600000},
				},
			},
		},
	}
	if TotalTeamPoints(&team) != 650021 {
		t.Fail()
	}
}

func TestTeamDisabled(t *testing.T) {
	team := report.SimpleTeam{
		Pause: true,
		Hosts: map[uuid.UUID]*report.SimpleHost{
			uuid.Must(uuid.NewV4()): {
				Pause: false,
				Services: map[uuid.UUID]*report.SimpleService{
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 1},
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 20},
				},
			},
			uuid.Must(uuid.NewV4()): {
				Pause: true,
				Services: map[uuid.UUID]*report.SimpleService{
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 300},
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 4000},
				},
			},
			uuid.Must(uuid.NewV4()): {
				Pause: false,
				Services: map[uuid.UUID]*report.SimpleService{
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 50000},
					uuid.Must(uuid.NewV4()): {Pause: false, Points: 600000},
				},
			},
		},
	}
	if TotalTeamPoints(&team) != 0 {
		t.Fail()
	}
}

func TestTeamNumber(t *testing.T) {
	team := report.SimpleTeam{Name: "Test 123 Team"}
	number, err := TeamNumber(&team)
	if err != nil {
		t.Error(err)
	}
	if number != 123 {
		t.Fail()
	}
}
