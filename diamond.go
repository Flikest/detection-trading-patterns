package detectedtradingpatterns

import mathutils "github.com/Flikest/detection-trading-patterns/math_utils"

func (p *patterns) Diamond(request request) PatternResult {
	response := PatternResult{}

	avgPoints := mathutils.CalcAveragesOHLC(request.data)

	zigzag := Zigzag(avgPoints, request.threshold, true)

}
