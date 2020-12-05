package report

import "testing"

func TestHostEnabled(t *testing.T) {
	h := Host{
		Enabled: true,
		Services: map[string]*Service{
			"a": {Enabled: true, Points: 1},
			"b": {Enabled: true, Points: 20},
			"c": {Enabled: true, Points: 300},
			"d": {Enabled: false, Points: 4000},
			"e": {Enabled: true, Points: 50000},
		},
	}

	if h.TotalPoints() != 50321 {
		t.Fail()
	}
}

func TestHostDisabled(t *testing.T) {
	h := Host{
		Enabled: false,
		Services: map[string]*Service{
			"a": {Enabled: true, Points: 1},
			"b": {Enabled: true, Points: 20},
			"c": {Enabled: true, Points: 300},
			"d": {Enabled: false, Points: 4000},
			"e": {Enabled: true, Points: 50000},
		},
	}

	if h.TotalPoints() != 0 {
		t.Fail()
	}
}
