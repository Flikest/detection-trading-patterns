package detectedtradingpatterns

import "math"

// парамметр deviation лучше определять как средний показатель волатильности в процентах
// например: волатильность за месяц = 2%, тогда в параметр deviation мы записываем 0.02
func Zigzag(data []float64, deviation float64, isUptrend bool) []int {
	var pivotIndices []int

	if len(data) == 0 {
		return pivotIndices
	}

	lastPivotIndex := 0
	lastPivotPrice := data[0]

	for i := 1; i < len(data); i++ {
		change := math.Abs(data[i]-lastPivotPrice) / lastPivotPrice

		if isUptrend {
			if data[i] < lastPivotPrice && change >= deviation {
				pivotIndices = append(pivotIndices, lastPivotIndex)
				lastPivotIndex = i
				lastPivotPrice = data[i]
				isUptrend = false
			}
		} else {
			if data[i] > lastPivotPrice && change >= deviation {
				pivotIndices = append(pivotIndices, lastPivotIndex)
				lastPivotIndex = i
				lastPivotPrice = data[i]
				isUptrend = true
			}
		}
	}
	pivotIndices = append(pivotIndices, lastPivotIndex)

	return pivotIndices
}
