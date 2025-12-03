package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var path string
	flag.StringVar(&path, "f", "", "input file")
	flag.Parse()
	if path == "" {
		panic("missing input file")
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	nums := make([]int, 0)
	nums12 := make([]int, 0)
	sum := 0
	sum12 := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		chars := strings.Split(line, "")
		chosen := choose(2, chars)
		chosen12 := choose(12, chars)
		num, _ := strconv.Atoi(strings.Join(chosen, ""))
		num12, _ := strconv.Atoi(strings.Join(chosen12, ""))
		nums = append(nums, num)
		nums12 = append(nums12, num12)
		sum += num
		sum12 += num12
	}
	fmt.Println("nums:", nums, "sum:", sum)
	fmt.Println("nums12:", nums12, "sum12:", sum12)
}

func choose(num int, chars []string) []string {
	result := make([]string, num)
	lastFind := -1
	for i := range num {
		skipLast := num - i - 1
		largest := "0"
		for j := lastFind + 1; j < len(chars)-skipLast-1; j++ {
			if chars[j] > largest {
				largest = chars[j]
				result[i] = chars[j]
				lastFind = j
			}
			// fmt.Println("largest:", largest, "lastFind:", lastFind, "result:", result, "i:", i)
		}
	}
	// fmt.Println("")
	return result
}
