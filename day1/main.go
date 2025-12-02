package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Dial struct {
	pos             int
	timesLandedZero int
	timesPassedZero int
}

func NewDial() *Dial {
	return &Dial{pos: 50}
}

func (d *Dial) Move(dir string, steps int) {
	diff := steps % 100 // can be > 100
	if dir == "L" {
		d.pos -= diff
		if d.pos < 0 {
			d.pos += 100
		}
	} else {
		d.pos += diff
		if d.pos > 99 {
			d.pos -= 100
		}
	}
	if d.pos == 0 {
		d.timesLandedZero++
	}
}

func (d *Dial) Move2(dir string, steps int) {
	for range steps {
		if dir == "L" {
			d.pos--
			if d.pos == 0 {
				d.timesPassedZero++
			}
			if d.pos < 0 {
				d.pos += 100
			}
		} else {
			d.pos++
			if d.pos == 100 {
				d.timesPassedZero++
				d.pos = 0
			}
		}
	}
	if d.pos == 0 {
		d.timesLandedZero++
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
	dial := NewDial()
	for {
		line, err := r.ReadString('\n')
		if line == "" {
			break
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dir := line[0]
		steps, err := strconv.Atoi(line[1 : len(line)-1])
		if err != nil {
			panic(err)
		}
		// dial.Move(string(dir), steps)
		dial.Move2(string(dir), steps)
	}
	//Number of times dial landed at zero: 1052
	//Number of times dial passed zero: 6295

	fmt.Printf("Number of times dial landed at zero: %d\n", dial.timesLandedZero)
	fmt.Printf("Number of times dial passed zero: %d\n", dial.timesPassedZero)
}
