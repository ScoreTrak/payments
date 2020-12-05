package report

import "testing"

func TestTeamEnabled(t *testing.T) {
	team := Team{
		Enabled: true,
		Hosts: map[string]*Host{
			"host-a": {
				Enabled: true,
				Services: map[string]*Service{
					"service-a": {Enabled: true, Points: 1},
					"service-b": {Enabled: true, Points: 20},
				},
			},
			"host-b": {
				Enabled: false,
				Services: map[string]*Service{
					"service-c": {Enabled: true, Points: 300},
					"service-d": {Enabled: true, Points: 4000},
				},
			},
			"host-c": {
				Enabled: true,
				Services: map[string]*Service{
					"service-e": {Enabled: true, Points: 50000},
					"service-f": {Enabled: true, Points: 600000},
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
		Enabled: false,
		Hosts: map[string]*Host{
			"host-a": {
				Enabled: true,
				Services: map[string]*Service{
					"service-a": {Enabled: true, Points: 1},
					"service-b": {Enabled: true, Points: 20},
				},
			},
			"host-b": {
				Enabled: false,
				Services: map[string]*Service{
					"service-c": {Enabled: true, Points: 300},
					"service-d": {Enabled: true, Points: 4000},
				},
			},
			"host-c": {
				Enabled: true,
				Services: map[string]*Service{
					"service-e": {Enabled: true, Points: 50000},
					"service-f": {Enabled: true, Points: 600000},
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
