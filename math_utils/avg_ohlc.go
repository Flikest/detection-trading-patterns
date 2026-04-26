package mathutils

import detectedtradingpatterns "github.com/Flikest/detection-trading-patterns"

func CalcAveragesOHLC(data []detectedtradingpatterns.OHLC) []float64 {
	averages := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		d := data[i]
		averages = append(averages, (d.open+d.high+d.low+d.close)/4.0)
	}
	return averages
}
