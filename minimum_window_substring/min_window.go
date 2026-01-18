package minimumwindowsubstring

import "math"

// MinWindow returns the smallest substring of s that contains all characters of t.
//
// Unicode requirement:
// - Treat s and t as sequences of runes (Unicode code points), not bytes.
// - Duplicates in t matter (frequency counts).
//
// Rules:
//   - Return the minimum-length contiguous substring of s (by rune positions)
//     that contains every rune in t with at least the same multiplicity.
//   - If no such substring exists, return "".
//   - If t is empty, return "".
func MinWindow(s, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	sr := []rune(s)
	tr := []rune(t)
	if len(tr) > len(sr) {
		return ""
	}

	need := make(map[rune]int, len(tr))
	for _, r := range tr {
		need[r]++
	}

	missing := len(tr)

	bestStart := 0
	bestLen := math.MaxInt

	left := 0
	for right := range sr {
		rc := sr[right]

		if need[rc] > 0 {
			missing--
		}
		need[rc]--

		for missing == 0 {
			windowLen := right - left + 1
			if windowLen < bestLen {
				bestLen = windowLen
				bestStart = left
			}

			lc := sr[left]
			need[lc]++
			if need[lc] > 0 {
				missing++
			}
			left++
		}
	}

	if bestLen == math.MaxInt {
		return ""
	}

	return string(sr[bestStart : bestStart+bestLen])
}
