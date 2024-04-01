package main

import (
	"fmt"
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

	matrix := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}

	fmt.Println(dynamicArray(6, matrix))
}
