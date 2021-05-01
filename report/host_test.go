package report

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/gofrs/uuid"
	"testing"
)

func TestHostEnabled(t *testing.T) {
	h := report.SimpleHost{
		Pause: false,
		Services: map[uuid.UUID]*report.SimpleService{
			uuid.Must(uuid.NewV4()): {Pause: false, Points: 1},
			uuid.Must(uuid.NewV4()): {Pause: false, Points: 20},
			uuid.Must(uuid.NewV4()): {Pause: false, Points: 300},
			uuid.Must(uuid.NewV4()): {Pause: true, Points: 4000},
			uuid.Must(uuid.NewV4()): {Pause: false, Points: 50000},
		},
	}

	if TotalHostPoints(&h) != 50321 {
		t.Fail()
	}
}

func TestHostDisabled(t *testing.T) {
	h := report.SimpleHost{
		Pause: true,
		Services: map[uuid.UUID]*report.SimpleService{
			uuid.Must(uuid.NewV4()): {Pause: false, Points: 1},
			uuid.Must(uuid.NewV4()): {Pause: false, Points: 20},
			uuid.Must(uuid.NewV4()): {Pause: false, Points: 300},
			uuid.Must(uuid.NewV4()): {Pause: true, Points: 4000},
			uuid.Must(uuid.NewV4()): {Pause: false, Points: 50000},
		},
	}

	if TotalHostPoints(&h) != 0 {
		t.Fail()
	}
}
