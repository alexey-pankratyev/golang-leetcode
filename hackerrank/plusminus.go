package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
	p := 0
	n := 0
	z := 0
	for i := range arr {
		switch{
			case math.Signbit(float64(arr[i])):
				n++
			case !math.Signbit(float64(arr[i])) && arr[i] != 0:
				p++
			case arr[i] == 0:
				z++  
		}				
	}
	fmt.Printf("%.6f\n",float64(p)/float64(len(arr)))
	fmt.Printf("%.6f\n",float64(n)/float64(len(arr)))
	fmt.Printf("%.6f\n",float64(z)/float64(len(arr)))
}

func main() {
	arrTemp := []int32{-4,3,-9,0,4,1}
    plusMinus(arrTemp)
}