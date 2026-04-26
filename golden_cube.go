package detectedtradingpatterns

import (
	mathutils "github.com/Flikest/detection-trading-patterns/math_utils"
)

func golden_cube(request request, candleWidth float64) []PatternResult {
	if len(request.data) < 4 {
		return nil
	}

	var response []PatternResult

	cubeWidth := candleWidth * 4

	leftDeviation := cubeWidth - cubeWidth*request.threshold
	rightDeviation := cubeWidth + cubeWidth*request.threshold

	for i := 3; i < len(request.data)-3; i++ {
		candle1Height := request.data[i-3].High - request.data[i-3].Low
		candle2Height := request.data[i-2].High - request.data[i-2].Low
		candle3Height := request.data[i-1].High - request.data[i-1].Low
		candle4Height := request.data[i].High - request.data[i].Low

		maxHeightCandle := max(candle1Height, candle2Height, candle3Height, candle4Height)
		minHeightCandle := min(candle1Height, candle2Height, candle3Height, candle4Height)

		isMinCandleFit := leftDeviation <= minHeightCandle && minHeightCandle <= rightDeviation
		isMaxCandleFit := leftDeviation <= maxHeightCandle && maxHeightCandle <= rightDeviation

		if isMaxCandleFit && isMinCandleFit {
			response = append(response, PatternResult{
				PatternName: "golden cube",
				Points: []PatternPoint{
					{Index: i - 3, Price: request.data[i-3].High},
					{Index: i - 2, Price: request.data[i-2].High},
					{Index: i - 1, Price: request.data[i-1].High},
					{Index: i, Price: request.data[i].High},
				},
				Confidence: mathutils.CalcConfidence(candle1Height+candle2Height+candle3Height+candle4Height, candleWidth*4, request.threshold),
			})
		} else {
			continue
		}
	}

	return response

}
