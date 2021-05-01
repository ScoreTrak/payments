package report

import "testing"

func TestHostEnabled(t *testing.T) {
	h := Host{
		Pause: true,
		Services: map[string]*Service{
			"a": {Pause: true, Points: 1},
			"b": {Pause: true, Points: 20},
			"c": {Pause: true, Points: 300},
			"d": {Pause: false, Points: 4000},
			"e": {Pause: true, Points: 50000},
		},
	}

	if h.TotalPoints() != 50321 {
		t.Fail()
	}
}

func TestHostDisabled(t *testing.T) {
	h := Host{
		Pause: false,
		Services: map[string]*Service{
			"a": {Pause: true, Points: 1},
			"b": {Pause: true, Points: 20},
			"c": {Pause: true, Points: 300},
			"d": {Pause: false, Points: 4000},
			"e": {Pause: true, Points: 50000},
		},
	}

	if h.TotalPoints() != 0 {
		t.Fail()
	}
}
