package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func diagonalDifference(arr [][]int32) int32 {
	var leftToRight int32
	var rightToLeft int32

	matrixLength := len(arr)

	fmt.Println(matrixLength)
	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			if i == j {
				leftToRight += arr[i][j]
			}

			fmt.Println(j, matrixLength, i, matrixLength-i-1)

			if j == matrixLength-i-1 {
				rightToLeft += arr[i][j]
			}
		}
	}

	//for i, h := range arr {
	//	for j, cell := range h {
	//		if i == j {
	//			leftToRight += cell
	//		}
	//	}
	//}
	//
	//for i, h := range arr {
	//	reverseArray := make([]int32, 0, len(h))
	//	for index := len(h) - 1; index >= 0; index-- {
	//		reverseArray = append(reverseArray, h[index])
	//	}
	//	for j, cell := range reverseArray {
	//		if i == j {
	//			rightToLeft += cell
	//		}
	//	}
	//}

	if result := leftToRight - rightToLeft; result > 1 {
		return result
	} else {
		return result * -1
	}
}

func plusMinus(arr []int32) {
	var negative, positive, zero float64
	var qtyNegative, qtyPositive, qtyZero float64
	arrayLength := float64(len(arr))

	for i := 0; i < int(arrayLength); i++ {
		if arr[i] < 0 {
			qtyNegative += 1
		} else if arr[i] > 0 {
			qtyPositive += 1
		} else {
			qtyZero += 1
		}
	}

	positive = qtyPositive / arrayLength
	negative = qtyNegative / arrayLength
	zero = qtyZero / arrayLength

	fmt.Printf("%.6f\n", positive)
	fmt.Printf("%.6f\n", negative)
	fmt.Printf("%.6f\n", zero)
}

func staircase(n int32) {
	intN := int(n)
	for i := 1; i <= intN; i++ {
		format := "%" + strconv.Itoa(intN) + "v"
		fmt.Println(fmt.Sprintf(format, strings.Repeat("#", i)))
	}
}

func reverseArray(a []int32) []int32 {
	reverse := make([]int32, 0, len(a))
	for index := len(a) - 1; index >= 0; index-- {
		reverse = append(reverse, a[index])
	}

	return reverse
}

//func hourglassSum(arr [][]int32) int32 {
//	maxSum := math.Inf(-1)
//
//	for line := 0; line < 4; line++ {
//		for column := 0; column < 4; column++ {
//			top := arr[line][column] + arr[line][column+1] + arr[line][column+2]
//			mid := arr[line+1][column+1]
//			bottom := arr[line+2][column] + arr[line+2][column+1] + arr[line+2][column+2]
//			sum := float64(top + mid + bottom)
//
//			if sum > maxSum {
//				maxSum = sum
//			}
//		}
//	}
//
//	return int32(maxSum)
//}

func dynamicArray(n int32, queries [][]int32) []int32 {
	var lastAns int32
	seqList := make([][]int32, 0)
	var result []int32
	for i := 0; i < int(n); i++ {
		seqList = append(seqList, make([]int32, 0))
	}

	for _, arr := range queries {
		x := arr[0]
		seq := (arr[1] ^ lastAns) % n
		size := int32(len(seqList[seq]))
		if x == 1 {
			seqList[seq] = append(seqList[seq], arr[2])
		} else if x == 2 {
			lastAns = seqList[seq][arr[2]%size]
			result = append(result, lastAns)
		}
	}
	return result
}

func rotateLeft(d int32, arr []int32) []int32 {
	tempArr := arr

	for i := 0; i < int(d); i++ {
		firstItem := tempArr[0]
		tempArr = append(tempArr[1:], firstItem)
	}

	return tempArr
}

func matchingStrings(stringList []string, queries []string) []int32 {
	results := make([]int32, len(queries))

	for _, str := range stringList {
		for index, query := range queries {
			if query == str {
				results[index] += 1
			}
		}
	}

	return results
}

func miniMaxSum(arr []int32) {
	// MINHA SOLUÇÃO ANTES DE PESQUISAR

	//var mini, maxSum float64
	//var bigger int32
	//var smaller int32 = math.MaxInt32
	//arrCopy := make([]int32, len(arr))
	//copy(arrCopy, arr)
	//
	//arrMini, arrMax := make([]int32, 0), make([]int32, 0)
	//
	//var biggerIndex, smallerIndex int
	//lastIndex := len(arr) - 1
	//
	//for index, num := range arr {
	//	if num > bigger {
	//		bigger = num
	//		biggerIndex = index
	//	}
	//
	//	if smaller > num {
	//		smaller = num
	//		smallerIndex = index
	//	}
	//}
	//
	//if biggerIndex == lastIndex {
	//	arrMini = append([]int32{}, arrCopy[:biggerIndex]...)
	//} else {
	//	arrMini = append(arrCopy[:biggerIndex], arrCopy[biggerIndex+1:]...)
	//}
	//
	//for _, num := range arrMini {
	//	mini += float64(num)
	//}
	//
	//if smallerIndex == lastIndex {
	//	arrMax = append([]int32{}, arr[:smallerIndex]...)
	//} else {
	//	arrMax = append(arr[:smallerIndex], arr[smallerIndex+1:]...)
	//}
	//
	//for _, num := range arrMax {
	//	maxSum += float64(num)
	//}
	//
	//fmt.Printf("%0.f %0.f", mini, maxSum)

	// SOLUÇÃO FINAL
	var mini, maxSum int64

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr)-1; i++ {
		mini += int64(arr[i])
		maxSum += int64(arr[i+1])
	}

	fmt.Println(mini, maxSum)

}

func birthdayCakeCandles(candles []int32) int32 {
	sort.Slice(candles, func(i, j int) bool {
		return candles[j] < candles[i]
	})

	tallestCandle := candles[0]
	var total int32

	for _, candle := range candles {
		if candle == tallestCandle {
			total += 1
		}
	}

	return total
}

func timeConversion(s string) string {
	format := s[8:]
	hours, minutes, seconds := s[:2], s[3:5], s[6:8]

	if hours == "12" && format == "AM" {
		hours = "00"
	} else if hours != "12" && format == "PM" {
		parse, err := strconv.Atoi(hours)
		if err != nil {
			panic(err)
		}
		hours = strconv.Itoa(parse + 12)
	}
	fmt.Println(minutes, seconds)
	return fmt.Sprintf("%s:%s:%s", hours, minutes, seconds)
}

func main() {
	//arrDiagonalDiff := [][]int32{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{9, 8, 9},
	//}
	//
	//fmt.Println(diagonalDifference(arrDiagonalDiff))

	//arrPlusMinus := []int32{-4, 3, -9, 0, 4, 1}
	//plusMinus(arrPlusMinus)

	//staircase(8)

	//arrReverseArray := []int32{1, 4, 3, 2}
	//fmt.Println(reverseArray(arrReverseArray))

	//matrix := [][]int32{
	//	{1, 1, 1, 0, 0, 0},
	//	{0, 1, 0, 0, 0, 0},
	//	{1, 1, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0},
	//}
	//
	//fmt.Println(hourglassSum(matrix))

	//matrix := [][]int32{
	//	{1, 0, 5},
	//	{1, 1, 7},
	//	{1, 0, 3},
	//	{2, 1, 0},
	//	{2, 1, 1},
	//}
	//
	//fmt.Println(dynamicArray(6, matrix))

	//arr := []int32{
	//	1, 2, 3, 4, 5,
	//}
	//
	//fmt.Println(rotateLeft(54, arr))

	//stringList := []string{"aba", "baba", "aba", "xzxb"}
	//queries := []string{"aba", "xzxb", "ab"}
	//
	//fmt.Println(matchingStrings(stringList, queries))

	//arr := []int32{396285104, 573261094, 759641832, 819230764, 364801279}
	//
	//miniMaxSum(arr)

	//arr := []int32{3, 2, 1, 3}
	//
	//fmt.Println(birthdayCakeCandles(arr))

	str := "08:01:00PM"
	fmt.Println(timeConversion(str))
}
