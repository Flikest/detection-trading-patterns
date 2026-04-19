package detectedtradingpatterns

import (
	"math"
)

type PatternPoint struct {
	Index int
	Price float64
}

type PatternResult struct {
	PatternName string
	Points      []PatternPoint
}

type request struct {
	data      []OHLC
	threshold float64
}

type avgOHLC struct {
	average float64
}

func findLocalMaximumAvg(data []float64, threshold float64) []PatternPoint {
	localMaximum := []PatternPoint{}

	for i := 1; i < len(data)-1; i++ {
		prevDiff := math.Abs(data[i] - data[i-1])
		nextDiff := math.Abs(data[i] - data[i+1])

		if data[i] > data[i-1] && data[i] > data[i+1] &&
			prevDiff > threshold && nextDiff > threshold {
			localMaximum = append(localMaximum, PatternPoint{
				Index: i,
				Price: data[i],
			})
		}
	}

	return localMaximum
}

func findLocalMininumAvg(data []float64, threshold float64) []PatternPoint {
	localMinimum := []PatternPoint{}

	for i := 1; i < len(data)-1; i++ {
		prevDiff := math.Abs(data[i] - data[i-1])
		nextDiff := math.Abs(data[i] - data[i+1])

		if data[i] < data[i-1] && data[i] < data[i+1] &&
			(prevDiff > threshold) &&
			(nextDiff > threshold) {
			localMinimum = append(localMinimum, PatternPoint{
				Index: i,
				Price: data[i],
			})
		}
	}

	return localMinimum
}

func calcAverages(data []OHLC) []float64 {
	averages := make([]float64, len(data))
	for _, d := range data {
		averages = append(averages, (d.open+d.high+d.low+d.close)/4.0)
	}
	return averages
}

// парамметр deviation лучше определять как средний показатель волатильности в процентах
// например: волатильность за месяц = 2%, тогда в параметр deviation мы записываем 0.02
func calculateZigZag(data []float64, deviation float64) []int {
	var pivotIndices []int

	if len(data) == 0 {
		return pivotIndices
	}

	lastPivotIndex := 0
	lastPivotPrice := data[0]

	isUptrend := true

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

func (p *Patterns) HeadAndShoulders(request request, deviation float64) PatternResult {
	response := PatternResult{}

	avgPoints := calcAverages(request.data)

	zigZag := calculateZigZag(avgPoints, deviation)

	localMinimums := findLocalMininumAvg(avgPoints, request.threshold)
	localMaximums := findLocalMaximumAvg(avgPoints, request.threshold)

	result := append(localMinimums, localMaximums...)

	if len(localMinimums) == 2 && len(localMaximums) == 3 {
		response.PatternName = "head and shoulders"
	} else if len(localMinimums) > 2 && len(localMaximums) > 3 {
		response.PatternName = "complex head and shoulders"
	} else if len(localMinimums) == 3 && len(localMaximums) == 2 {
		response.PatternName = "reversed head and shoulders"
	} else if len(localMinimums) > 3 && len(localMaximums) > 2 {
		response.PatternName = "reversed head and shoulders"
	}

	response.Points = result

	return response
}
