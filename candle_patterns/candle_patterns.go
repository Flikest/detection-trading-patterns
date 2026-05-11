package candlepatterns

type CandlePatternPoint struct {
	Index uint
	Price float64
}

type CandlePatterns struct{}

type CandlePatternResult struct {
	PatternName   string
	PatternPoints []CandlePatternPoint
}
