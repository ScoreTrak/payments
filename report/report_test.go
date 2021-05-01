package report

import "reflect"
import "testing"

func TestReport(t *testing.T) {
	report := Report{
		Round: 123,
		Teams: map[string]*Team{
			"team-1": {
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
			},
		},
	}
	expected := map[int]uint{123: 650021}
	actual := report.PointsPerTeam()
	if !reflect.DeepEqual(expected, actual) {
		t.Fail()
	}
}
