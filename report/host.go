package report

type Host struct {
	Enabled  bool
	Services map[string]*Service
}

func (h *Host) TotalPoints() uint {
	if !h.Enabled {
		return 0
	}
	var total uint = 0
	for _, service := range h.Services {
		total += service.TotalPoints()
	}
	return total
}
