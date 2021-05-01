package report

import "github.com/ScoreTrak/ScoreTrak/pkg/report"

func TotalHostPoints(h *report.SimpleHost) uint64 {
	if h.Pause {
		return 0
	}
	var total uint64 = 0
	for _, service := range h.Services {
		total += TotalServicePoints(service)
	}
	return total
}
