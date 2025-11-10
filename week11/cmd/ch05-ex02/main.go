// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("meetweight.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for i, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("평균 고기 소모량: %0.2f\n", sum/sampleCount)
}
