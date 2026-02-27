package lib

import "backend/internal/domain"


func CalculateSummary(data []domain.Payment) map[string]int {
	summary := map[string]int{
		"total":   len(data),
		"success": 0,
		"failed":  0,
	}

	for _, p := range data {
		switch p.Status {
		case "success":
			summary["success"]++
		case "failed":
			summary["failed"]++
		}
	}

	return summary
}