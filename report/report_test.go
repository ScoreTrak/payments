package report

import (
	"fmt"
	report2 "github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/gofrs/uuid"
	"reflect"
)
import "testing"

func TestReport(t *testing.T) {
	report := report2.SimpleReport{
		Round: 123,
		Teams: map[uuid.UUID]*report2.SimpleTeam{
			uuid.Must(uuid.NewV4()): {
				Name:  "team-123",
				Pause: false,
				Hosts: map[uuid.UUID]*report2.SimpleHost{
					uuid.Must(uuid.NewV4()): {
						Pause: false,
						Services: map[uuid.UUID]*report2.SimpleService{
							uuid.Must(uuid.NewV4()): {Pause: false, Points: 1},
							uuid.Must(uuid.NewV4()): {Pause: false, Points: 20},
						},
					},
					uuid.Must(uuid.NewV4()): {
						Pause: true,
						Services: map[uuid.UUID]*report2.SimpleService{
							uuid.Must(uuid.NewV4()): {Pause: false, Points: 300},
							uuid.Must(uuid.NewV4()): {Pause: false, Points: 4000},
						},
					},
					uuid.Must(uuid.NewV4()): {
						Pause: false,
						Services: map[uuid.UUID]*report2.SimpleService{
							uuid.Must(uuid.NewV4()): {Pause: false, Points: 50000},
							uuid.Must(uuid.NewV4()): {Pause: false, Points: 600000},
						},
					},
				},
			},
		},
	}
	expected := map[int]uint64{123: 650021}
	actual := PointsPerTeam(&report)
	fmt.Println(actual, expected)
	if !reflect.DeepEqual(expected, actual) {
		t.Fail()
	}
}
