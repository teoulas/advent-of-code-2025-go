package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) Includes(x int) bool {
	return x >= r.Start && x <= r.End
}

type Database struct {
	Ranges      []Range
	Ingredients []int
}

func (db *Database) NumFresh() int {
	num := 0
	for _, ing := range db.Ingredients {
		for _, r := range db.Ranges {
			if r.Includes(ing) {
				num++
				break
			}
		}
	}
	return num
}

func (db *Database) NumAllFresh() int {
	// Sort ranges by start position
	ranges := make([]Range, len(db.Ranges))
	copy(ranges, db.Ranges)
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.Start - b.Start
	})

	// Merge overlapping ranges and count
	total := ranges[0].End - ranges[0].Start + 1
	currentEnd := ranges[0].End

	for i := 1; i < len(ranges); i++ {
		if ranges[i].Start <= currentEnd+1 {
			// Ranges overlap or are adjacent
			if ranges[i].End > currentEnd {
				total += ranges[i].End - currentEnd
				currentEnd = ranges[i].End
			}
		} else {
			// New separate range
			total += ranges[i].End - ranges[i].Start + 1
			currentEnd = ranges[i].End
		}
	}

	return total
}

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
	firstPart := true
	db := Database{}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			firstPart = false
			continue
		}
		if firstPart {
			strs := strings.SplitN(line, "-", 2)
			start, err := strconv.Atoi(strs[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(strs[1])
			if err != nil {
				panic(err)
			}
			r := Range{Start: start, End: end}
			db.Ranges = append(db.Ranges, r)
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			db.Ingredients = append(db.Ingredients, i)
		}
	}
	fmt.Printf("Number of fresh ingredients: %d\n", db.NumFresh())
	fmt.Printf("Number of spoiled ingredients: %d\n", db.NumAllFresh())
}
