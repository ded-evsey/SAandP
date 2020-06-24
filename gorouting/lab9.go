package main

import "fmt"

func maxValueInArray(array []int, maxValueRes chan int) {
	maxValue := 0
	for _,value := range array{
		if 0 < value && maxValue < value{
			maxValue = value
		}
	}
	maxValueRes <- maxValue
}

func main()  {
	array := []int{3,12,10,20,156,8,4,3,-5,0,100,-6,8,145,-1}
	c := make(chan int)

	go maxValueInArray(array[0:4], c)
	go maxValueInArray(array[5:9], c)
	go maxValueInArray(array[10:14], c)
	maxArray := []int{<-c,<-c,<-c}
	cLast := make(chan  int)
	go maxValueInArray(maxArray, cLast)
	max := <-cLast
	fmt.Print(max)

}