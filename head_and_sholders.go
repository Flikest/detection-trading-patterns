package detectedtradingpatterns

import mathutils "github.com/Flikest/detection-trading-patterns/math_utils"

func (p *patterns) HeadAndShoulders(request request, deviation float64) PatternResult {
	response := PatternResult{}

	avgPoints := mathutils.CalcAveragesOHLC(request.data)

	zigzag := Zigzag(avgPoints, deviation, true)

	znachimieTochki := []float64{}

	for _, point := range zigzag {
		znachimieTochki = append(znachimieTochki, avgPoints[point])
	}

	localMinimums := findLocalMininumsAvg(znachimieTochki, request.threshold)
	localMaximums := findLocalMaximumsAvg(znachimieTochki, request.threshold)

	result := append(localMinimums, localMaximums...)

	if len(localMinimums) == 2 && len(localMaximums) == 3 &&
		localMaximums[0].Price+localMaximums[2].Price <= max(localMaximums[0].Price, localMaximums[2].Price)*deviation &&
		max(localMaximums[0].Price, localMaximums[1].Price, localMaximums[2].Price) == localMaximums[1].Price {
		response.PatternName = "head and shoulders"
	} else if len(localMinimums) > 2 && len(localMaximums) > 3 {
		response.PatternName = "complex head and shoulders"
	} else if len(localMinimums) == 3 && len(localMaximums) == 2 &&
		localMinimums[0].Price+localMinimums[2].Price <= max(localMinimums[0].Price, localMinimums[2].Price)*deviation &&
		max(localMinimums[0].Price, localMinimums[1].Price, localMinimums[2].Price) == localMinimums[1].Price {
		response.PatternName = "reversed head and shoulders"
	} else if len(localMinimums) > 3 && len(localMaximums) > 2 {
		response.PatternName = "reversed complex head and shoulders"
	}

	response.Points = result

	return response
}
