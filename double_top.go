package detectedtradingpatterns

func (p *patterns) DoubleTop(data []float64, threshold float64) PatternResult {
	localMaximums := findLocalMaximumsAvg(data, threshold)
	localMinimums := findLocalMininumsAvg(data, threshold)

	response := PatternResult{}

	if len(localMaximums) == 2 && len(localMinimums) == 1 {
		maxTop := max(localMaximums[0].Price, localMaximums[1].Price)
		minTop := min(localMaximums[0].Price, localMaximums[1].Price)

		if maxTop-minTop <= maxTop/100*5 {
			response.PatternName = "double top"
		}
	} else if len(localMinimums) == 2 && len(localMaximums) == 1 {
		maxBottom := max(localMinimums[0].Price, localMinimums[1].Price)
		minBottom := min(localMinimums[0].Price, localMinimums[1].Price)

		if maxBottom-minBottom <= maxBottom/100*5 {
			response.PatternName = "double bottom"
		}
	}

	response.Points = append(localMaximums, localMinimums...)

	return response
}
