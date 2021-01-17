package main

import (
	"./xdata"
	"fmt"
	)

func main()  {
	arrayValues := []float64{122.2, 112,0.2,-15,-186}
	fmt.Println(xdata.MiddleValue(arrayValues))
	fmt.Println(xdata.MedianValue(arrayValues))
	fmt.Println(xdata.MinValue(arrayValues))
	fmt.Println(xdata.MaxValue(arrayValues))
	fmt.Println(xdata.MiddleGeoValue(arrayValues))
}