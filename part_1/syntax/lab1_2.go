package main
import "fmt"

// Реализовать функцию вычисления суммы элементов числового массива, переданного как параметр в эту функцию.
// Реализовать функцию вычисления суммы положительных и сумму отрицательных элементов массива, переданного как параметр в эту функцию.
func sums(listValue []float64)(float64, float64, float64)	{
	total := 0.0
	totalPositive := 0.0
	totalAppositive := 0.0
	for _, v := range listValue{
		total += v
		if v < 0 {
			totalAppositive += v
		} else {
			totalPositive += v
		}
	}
	return total, totalPositive, totalAppositive
}

func main()  {
	listValue := []float64{-0.2, 0.1, 12, -128}
	total, totalPositive, totalAppositive := sums(listValue)
	fmt.Println("Sum: ", total)
	fmt.Println("Sum positive: ", totalPositive)
	fmt.Println("Sum appositive:", totalAppositive)
}
