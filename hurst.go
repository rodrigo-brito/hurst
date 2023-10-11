package hurst

import (
	"math"

	"gonum.org/v1/gonum/stat"
)

func Estimate(price []float64, minLag, maxLag int) float64 {
	lags := make([]int, maxLag-minLag+1)
	tau := make([]float64, maxLag-minLag+1)

	for lag := minLag; lag <= maxLag; lag++ {
		lags[lag-minLag] = lag
		diff := make([]float64, len(price)-lag)
		for i := 0; i < len(diff); i++ {
			diff[i] = price[i+lag] - price[i]
		}
		stdDev := stat.StdDev(diff, nil)
		tau[lag-minLag] = stdDev
	}

	logLags := make([]float64, len(lags))
	logTau := make([]float64, len(tau))
	for i := 0; i < len(lags); i++ {
		logLags[i] = math.Log10(float64(lags[i]))
		logTau[i] = math.Log10(tau[i])
	}

	m, _ := linearRegression(logLags, logTau)
	return m
}

func linearRegression(x, y []float64) (float64, float64) {
	n := len(x)
	var sumX, sumY, sumXY, sumXSquare float64

	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
	}

	m := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumXSquare - sumX*sumX)
	return m, 0
}
