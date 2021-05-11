package main

import (
	"bufio"
	"fmt"
	"gotest/other"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) (string, string, string, error) {

	pos, neg, zero := 0, 0, 0

	for _, n := range arr {
		if n == 0 {
			zero++
		} else if n < 0 {
			neg++
		} else {
			pos++
		}
	}
	l := len(arr)

	s1 := fmt.Sprintf("%0.6f", float64(pos)/float64(l))
	s2 := fmt.Sprintf("%0.6f", float64(neg)/float64(l))
	s3 := fmt.Sprintf("%0.6f", float64(zero)/float64(l))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	return s1, s2, s3, nil
}

func main() {

	another.Dirdemo()
	arr := []int32{5,-6,7,5,-7,5,5}
	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func processInput() []int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32
	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}
	return arr
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
