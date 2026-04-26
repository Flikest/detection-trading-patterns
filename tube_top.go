package detectedtradingpatterns

import (
	"math"

	mathutils "github.com/Flikest/detection-trading-patterns/math_utils"
)

func tubeTop(data []OHLC, treshold float64) []PatternResult {
	if len(data) < 2 {
		return nil
	}

	var response []PatternResult
	for i := 0; i < len(data)-1; i++ {
		candle1 := data[i]
		candle2 := data[i+1]

		isBearishFirstCandle := candle1.Close < candle1.Open
		isBullishSecondCandle := candle2.Close > candle2.Open

		var enough bool

		if treshold == 0 {
			enough = true
		}

		diff := math.Abs(candle1.Low - candle2.Low)
		avg := mathutils.CalcAverages([]float64{candle1.Low, candle2.Low})

		if avg == 0 {
			enough = diff <= treshold
		} else {
			enough = (diff / avg) <= treshold
		}

		if isBearishFirstCandle && isBullishSecondCandle && enough {
			response = append(response, PatternResult{
				PatternName: "tube_Top",
				Points:      []PatternPoint{{Index: i, Price: candle1.High}, {Index: i + 1, Price: candle1.High}},
				Confidence:  mathutils.CalcConfidence(candle1.Low, candle2.Low, treshold),
			})
		} else {
			response = append(response, PatternResult{
				PatternName: "tube_bottom",
				Points:      []PatternPoint{{Index: i, Price: candle1.High}, {Index: i + 1, Price: candle1.High}},
				Confidence:  mathutils.CalcConfidence(candle1.Low, candle2.Low, treshold),
			})
		}
	}

	return response
}
