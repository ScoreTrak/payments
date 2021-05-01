package report

import "testing"

func TestServiceEnabled(t *testing.T) {
	s := Service{Pause: true, Points: 123}

	if s.TotalPoints() != 123 {
		t.Fail()
	}
}

func TestServiceDisabled(t *testing.T) {
	s := Service{Pause: false, Points: 123}

	if s.TotalPoints() != 0 {
		t.Fail()
	}
}
