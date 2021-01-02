package main

import (
	"error/error"
	"fmt"
	"math"
)

func circleArea0(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func circleArea1(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &errors.AreaError{Err: "radius is negative", Radius: radius}
	}
	return math.Pi * radius * radius, nil
}
func main() {
	radius := -20.0
	//area, err := circleArea0(radius)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//基于错误，反馈更多信息radius
	area, err := circleArea1(radius)
	if err != nil{
		if err,ok := err.(*errors.AreaError);ok{
			fmt.Printf("Radius %0.2f is less than zero\n", err.Radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}