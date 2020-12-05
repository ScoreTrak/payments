package report

import "reflect"
import "testing"

func TestReport(t *testing.T) {
	report := Report{
		Round: 123,
		Teams: map[string]*Team{
			"team-1": {
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
			},
		},
	}
	expected := map[int]uint{123: 650021}
	actual := report.PointsPerTeam()
	if !reflect.DeepEqual(expected, actual) {
		t.Fail()
	}
}
