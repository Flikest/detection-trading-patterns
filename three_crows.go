package detectedtradingpatterns

func threeCrows(data []OHLC, treshold float64) []PatternResult {
	if len(data) < 6 {
		return nil
	}
	
	var response []PatternResult

	for batch := 5; batch < len(data)-1; batch++ {
		candle1 := data[batch-5]
		candle2 := data[batch-4]
		candle3 := data[batch-3]
		candle4 := data[batch-2]
		candle5 := data[batch-1]
		candle6 := data[batch]

		isFirstThreeCandleBearish := candle1.Close < candle1.Open && candle2.Close > candle2.Open && candle3.Close > candle3.Open
		isSecondThreeCandleBullish := candle4.Open < candle4.Close && candle5.Open < candle5.Close && candle6.Open < candle6.Close

		if !isFirstThreeCandleBearish || !isSecondThreeCandleBullish {
			continue
		}

		lengthFirstThreeCandles := candle1.High + candle1.Low + candle2.High + candle2.Low + candle3.High + candle3.Low
		lengthSecondThreeCandles := candle4.High + candle4.Low + candle5.High + candle5.Low + candle6.High + candle6.Low
		if lengthFirstThreeCandles > lengthSecondThreeCandles {
			continue
		} else {
			response = append(response, PatternResult{
				PatternName: "three crows",
				Points: []PatternPoint{
					{Index: batch - 2, Price: candle4.High},
					{Index: batch - 1, Price: candle5.High},
					{Index: batch, Price: candle6.High},
				},
				Confidence: ,
			})
		}
	}

	return response
}
