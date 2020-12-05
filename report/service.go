package report

type Service struct {
	Name    string
	Enabled bool
	Points  uint
}

func (s *Service) TotalPoints() uint {
	if !s.Enabled {
		return 0
	}
	return s.Points
}
