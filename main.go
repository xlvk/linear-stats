package main

import (
	
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
	"strconv"
	"sort"
)

func main() {
	var FMe []int
	var arr []int
	sum := 0.0
	sumNum := 0.0
	// mean := 0
	// sd := 0.0
	avr := 0.0
	readFile, err := os.Open("data.txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		ha, e := strconv.Atoi(fileScanner.Text())
    	if e != nil {
        	fmt.Printf("%T \n %v", ha, ha)
		}
		arr = append(arr, ha)
		FMe = append(FMe, ha)
		sum = sum + float64(ha)
		sumNum = sumNum + float64(len(arr)-1)
	}

	sort.Ints(arr)
	c := FMe[0]
	m := float64(FMe[2] - FMe[0])/2
	// y := 0
	// fmt.Print("The Average is: ")
	avr = math.Round((sum-(sumNum*m))/float64(len(arr))*1000000)/1000000

	// fmt.Println(avr)
	// fmt.Print("The Median is: ")
	// if len(arr)%2 == 0 {
	// 	mean = int(math.Ceil(float64((arr[(len(arr)/2)-1]) + (arr[(len(arr)/2)]))/2))
	// } else {
	// 	mean = int(math.Ceil(float64(arr[(len(arr)/2)])))
	// }
	// // fmt.Println(mean)
	// // fmt.Print("The Variance is: ")
	// for j := 0; j < len(arr); j++ {
	// 	sd += math.Pow((float64(arr[j])- avr), 2)
	// }
	// varr := math.Round(sd / float64(len(arr)))
	// fmt.Println(int(varr))
	// fmt.Print("The Standard Deviation is: ")
	// sd = math.Round(math.Sqrt(sd / float64(len(arr))))
	// fmt.Println(sd)
	fmt.Println("Linear Regression Line: y =", avr, "x +", c)
		fmt.Println("Pearson Correlation Coefficient:" )
}
