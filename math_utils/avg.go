package mathutils

import "math"

func CalcAverages(data []float64) float64 {
	var sum float64 = 0
	for _, d := range data {
		sum = sum + math.Abs(d)
	}
	return sum / float64(len(data))
}
