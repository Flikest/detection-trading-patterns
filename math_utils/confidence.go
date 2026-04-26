package mathutils

import "math"

func CalcConfidence(a, b, treshold float64) float64 {
	if treshold == 0 {
		if a == b {
			return 1.0
		}
	}

	avg := CalcAverages([]float64{a, b})

	if avg == 0 {
		if math.Abs(a-b) <= treshold {
			return 1.0
		}
		return 0.0
	}

	deffPercent := math.Abs(a-b) / avg
	confidence := 1.0 - (deffPercent / treshold)

	return max(0, min(1, confidence))
}
