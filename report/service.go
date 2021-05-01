package report

type Service struct {
	Name   string
	Pause  bool
	Points uint
}

func (s *Service) TotalPoints() uint {
	if !s.Pause {
		return 0
	}
	return s.Points
}
