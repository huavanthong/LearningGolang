/***************************************************************


Reference:
	- https://www.hackerrank.com/challenges/2d-array/problem
***************************************************************/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

const MaxUint32 = ^uint32(0)
const MinUint32 = 0
const MaxInt32 = int32(MaxUint32 >> 1)
const MinInt32 = -MaxInt32 - 1

func hourglassSum(arr [][]int32) int32 {
	var offsetx = [2]int32{-1, 1}
	var offsety = [3]int32{-1, 0, 1}

	// Write your code here
	var max int32 = MinInt32
	var result int32 = 0

	for i := int32(1); i < 5; i++ {
		for j := int32(1); j < 5; j++ {
			result = arr[i][j]
			fmt.Printf("Get value at center: arr[%d][%d] = %d \n", i, j, result)
			fmt.Println("Starting to calculate value around....")
			for x := 0; x < len(offsetx); x++ {
				for y := 0; y < len(offsety); y++ {
					result += arr[i+offsetx[x]][j+offsety[y]]
					fmt.Println("Val: ", arr[i+offsetx[x]][j+offsety[y]])
					fmt.Println("Sum: ", result)
				}
			}

			fmt.Printf("Ending, compare result %d and max %d\n", result, max)
			if result > max {
				max = result
			}
		}
	}

	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
