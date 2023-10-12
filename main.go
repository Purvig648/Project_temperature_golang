package main

import (
	"assignments/3bProject/temperaturee"
	"fmt"
)

func main() {
	var temp float64
	fmt.Print("Enter the temperature ")
	fmt.Scan(&temp)
	i := temperaturee.Convert(temp)

	fmt.Println("The Fahrenheit value is: ", i)
}
