package mathutils

import detectedtradingpatterns "github.com/Flikest/detection-trading-patterns"

func BearishCandle(ohlc detectedtradingpatterns.OHLC) bool {
	if ohlc.Open > ohlc.Close {
		return true
	} else {
		return false
	}
}
