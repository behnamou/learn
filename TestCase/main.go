package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Mid(array []int) float64 {
	sum := 0
	for _, item := range array {
		sum += item
	}
	if len(array) == 0 {
		return 0
	}
	return float64(sum) / float64(len(array))
}

func PopulationVariance1(input string) float64 {
	var stringArray []string = strings.Split(input, " ")
	length := len(stringArray)
	var array []int = make([]int, length)
	var temp int

	for i := 0; i < len(stringArray); i++ {
		temp, _ = strconv.Atoi(stringArray[i])
		array[i] = temp
	}

	var mid float64 = Mid(array[:])
	var sum float64 = 0
	for _, item := range array {
		sum += math.Pow((float64(item) - mid), 2)
	}

	return sum / float64(len(array))
}

func PopulationVariance2(input string) float64 {
	array := strings.Split(input, " ")
	var power float64 = 0
	var counter float64 = 0
	var sum float64 = 0
	var strToFloat float64
	var strToInt int

	for _, item := range array {
		strToInt, _ = strconv.Atoi(item)
		strToFloat = float64(strToInt)
		sum += float64(strToFloat)
		power += math.Pow(strToFloat, 2)
		counter += 1
	}
	average := sum / counter
	return (power / counter) - (math.Pow(average, 2))

}

func selectType() int {
	var varianceType string
	var rVal int

	for {
		fmt.Print("Choose Type\n1-Population Variance 1\n2-Population Variance 2\n=> ")
		fmt.Scanln(&varianceType)
		if varianceType == "1" || varianceType == "2" {
			rVal, _ = strconv.Atoi(varianceType)
			return rVal
		} else {
			fmt.Println("Wrong!!! Please choose correctly")
		}
	}
}

func main() {
	var input string
	varianceType := selectType()
	fmt.Print("Enter numbers with space between them :\n=> ")
	in := bufio.NewReader(os.Stdin)
	input, _ = in.ReadString('\n')

	input = input[:len(input)-2]

	if varianceType == 1 {
		fmt.Println(PopulationVariance1(input))
	} else {
		fmt.Println(PopulationVariance2(input))
	}
}
