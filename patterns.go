package detectedtradingpatterns

import "math"

type patterns struct{}

type OHLC struct {
	Open, High, Low, Close float64
}

type PatternPoint struct {
	Index int
	Price float64
}

type PatternResult struct {
	PatternName string
	Points      []PatternPoint
	Confidence  float64
}

type request struct {
	data      []OHLC
	threshold float64
}

type avgOHLC struct {
	average float64
}

func findLocalMaximumsAvg(data []float64, threshold float64) []PatternPoint {
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

func findLocalMininumsAvg(data []float64, threshold float64) []PatternPoint {
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
