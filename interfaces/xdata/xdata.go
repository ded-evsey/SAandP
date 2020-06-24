package xdata

import (
	"math"
	"sort"
)
func MiddleValue(arrayValues []float64) float64 {
	sum := 0.0
	for _, value := range arrayValues{
		sum += value
	}
	return sum / float64(len(arrayValues))
}

func MinValue(arrayValues []float64) float64  {
	min := arrayValues[0]
	for _, value := range arrayValues[1:]{
		if value < min {
			min = value
		}
	}
	return min
}

func MaxValue(arrayValues []float64) float64  {
	max := arrayValues[0]
	for _, value := range arrayValues[1:]{
		if value > max {
			max = value
		}
	}
	return max
}

func MedianValue(arrayValues []float64) float64  {
	sort.Float64s(arrayValues)
	if len(arrayValues) % 2 == 0{
		return arrayValues[(len(arrayValues)+1)/2]
	}
	return (arrayValues[len(arrayValues) / 2] + arrayValues[len(arrayValues) / 2 + 1]) / 2
}

func MiddleGeoValue(arrayValues []float64) float64  {
	pos := arrayValues[0]
	for _, value := range arrayValues[1:]{
		pos *= value
	}
	return math.Pow(pos, float64(1/len(arrayValues)))
}