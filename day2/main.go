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
	inv := make([]int, 0)
	inv2 := make([]int, 0)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		for part := range strings.SplitSeq(line, ",") {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}

			bounds := strings.Split(part, "-")
			beg, _ := strconv.Atoi(bounds[0])
			end, _ := strconv.Atoi(bounds[1])
			inv = append(inv, invalidIds(beg, end)...)
			inv2 = append(inv2, invalidIds2(beg, end)...)
		}
	}
	sum := 0

	for _, id := range inv {
		sum += id
	}
	fmt.Printf("Sum of invalid IDs: %d\n", sum)

	sum2 := 0
	for _, id := range inv2 {
		sum2 += id
	}
	fmt.Printf("Sum of invalid IDs2: %d\n", sum2)

}

func invalidIds(a, b int) []int {
	invalids := make([]int, 0)
	for z := a; z <= b; z++ {
		str := fmt.Sprintf("%d", z)
		l := len(str)
		if l%2 != 0 {
			continue
		}
		if str[0:l/2] == str[l/2:] {
			fmt.Printf("Invalid: %d\n", z)
			invalids = append(invalids, z)
		}
	}
	return invalids
}

func invalidIds2(a, b int) []int {
	invalids := make([]int, 0)
	for z := a; z <= b; z++ {
		str := fmt.Sprintf("%d", z)
		if isInvalid(str) {
			fmt.Printf("Invalid: %d\n", z)
			invalids = append(invalids, z)
		}
	}
	return invalids
}

func isInvalid(s string) bool {
	n := len(s)
	if n < 2 {
		return false
	}
	for l := 1; l <= n/2; l++ {
		if n%l != 0 {
			continue
		}
		p := s[:l]
		if strings.Repeat(p, n/l) == s {
			return true
		}
	}
	return false
}
