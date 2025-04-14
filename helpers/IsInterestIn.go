package helpers

import "slices"

func IsInterestIn(interesets []int, i int) string {
	if slices.Contains(interesets, i) {
		return "✅"
	} else {
		return ""
	}
}
