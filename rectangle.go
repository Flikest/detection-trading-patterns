package detectedtradingpatterns

import (
	mathutils "github.com/Flikest/detection-trading-patterns/math_utils"
)

func reactangle(request request) []PatternResult {
	var response []PatternResult

	avgOHLC := mathutils.CalcAveragesOHLC(request.data)

	zigzag := Zigzag(avgOHLC, request.threshold, true)

	points := []float64{}

	for _, j := range zigzag {
		points = append(points, avgOHLC[j])
	}

	localMaximums := findLocalMaximumsAvg(points, request.threshold)
	localMinimums := findLocalMininumsAvg(points, request.threshold)

	for i := 0; i < len(localMaximums)-1; i++ {

	}

}
