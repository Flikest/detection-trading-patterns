package mathutils

import detectedtradingpatterns "github.com/Flikest/detection-trading-patterns"

func IsBullishCandle(ohlc detectedtradingpatterns.OHLC) bool {
	if ohlc.Close > ohlc.Open {
		return true
	} else {
		return false
	}
}
