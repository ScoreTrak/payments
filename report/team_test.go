package report

import "testing"

func TestTeamEnabled(t *testing.T) {
	team := Team{
		Pause: true,
		Hosts: map[string]*Host{
			"host-a": {
				Pause: true,
				Services: map[string]*Service{
					"service-a": {Pause: true, Points: 1},
					"service-b": {Pause: true, Points: 20},
				},
			},
			"host-b": {
				Pause: false,
				Services: map[string]*Service{
					"service-c": {Pause: true, Points: 300},
					"service-d": {Pause: true, Points: 4000},
				},
			},
			"host-c": {
				Pause: true,
				Services: map[string]*Service{
					"service-e": {Pause: true, Points: 50000},
					"service-f": {Pause: true, Points: 600000},
				},
			},
		},
	}
	if team.TotalPoints() != 650021 {
		t.Fail()
	}
}

func TestTeamDisabled(t *testing.T) {
	team := Team{
		Pause: false,
		Hosts: map[string]*Host{
			"host-a": {
				Pause: true,
				Services: map[string]*Service{
					"service-a": {Pause: true, Points: 1},
					"service-b": {Pause: true, Points: 20},
				},
			},
			"host-b": {
				Pause: false,
				Services: map[string]*Service{
					"service-c": {Pause: true, Points: 300},
					"service-d": {Pause: true, Points: 4000},
				},
			},
			"host-c": {
				Pause: true,
				Services: map[string]*Service{
					"service-e": {Pause: true, Points: 50000},
					"service-f": {Pause: true, Points: 600000},
				},
			},
		},
	}
	if team.TotalPoints() != 0 {
		t.Fail()
	}
}

func TestTeamNumber(t *testing.T) {
	team := Team{Name: "Test 123 Team"}
	number, err := team.Number()
	if err != nil {
		t.Error(err)
	}
	if number != 123 {
		t.Fail()
	}
}
