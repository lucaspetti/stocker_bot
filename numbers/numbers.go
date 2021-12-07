package numbers

import (
	"fmt"
	"math"
)

var sizes = []string{"-", "K", "M", "B", "T"}

// FormatSuffix formats big numbers with corresponding suffix
func FormatSuffix(n int64) string {
	base := 1000.0
	number := uint64(n)

	if number < 1000 {
		return fmt.Sprintf("%d", number)
	}
	lognum := math.Log(float64(number)) / math.Log(base)

	e := math.Floor(lognum)
	val := math.Floor(float64(number)/math.Pow(base, e)*10) / 10
	suffix := sizes[int(e)]

	return fmt.Sprintf("%.2f%s", val, suffix)
}
