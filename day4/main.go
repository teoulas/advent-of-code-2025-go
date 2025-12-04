package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Map struct {
	Rows []string
}

func (m *Map) Dump() {
	for _, row := range m.Rows {
		fmt.Println(row)
	}
}

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func (m *Map) NumAccessible() (int, [][]int) {
	count := 0
	cellsToRemove := make([][]int, 0)
	for i := range len(m.Rows) {
		for j := range len(m.Rows[i]) {
			if m.Rows[i][j] != '@' {
				continue
			}
			rollsAround := 0
			for _, dir := range directions {
				x := i + dir[0]
				y := j + dir[1]
				if x < 0 || x >= len(m.Rows) || y < 0 || y >= len(m.Rows[x]) {
					continue
				}
				if m.Rows[x][y] == '.' {
					continue
				}
				rollsAround++
			}
			if rollsAround < 4 {
				count++
				cellsToRemove = append(cellsToRemove, []int{i, j})
			}
		}
	}
	return count, cellsToRemove
}

func (m *Map) RemoveCells(cellsToRemove [][]int) {
	for _, cell := range cellsToRemove {
		m.Rows[cell[0]] = m.Rows[cell[0]][:cell[1]] + "." + m.Rows[cell[0]][cell[1]+1:]
	}
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
	m := &Map{}
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
			continue
		}
		m.Rows = append(m.Rows, line)
	}
	total := 0
	i := 0
	for {
		fmt.Print("Iteration ", i+1, ": ")
		count, cellsToRemove := m.NumAccessible()
		fmt.Println(count)
		if count == 0 {
			break
		}
		total += count
		m.RemoveCells(cellsToRemove)
		i++
	}
	fmt.Println("Total removed:", total)
}
