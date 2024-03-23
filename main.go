package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	inputFile = "./measurements.txt"
)

type Measurement struct {
	Sum   float64
	Count int
	Max   float64
	Min   float64
}

func main() {
	start := time.Now()
	measurementsFile, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	defer measurementsFile.Close()

	scanner := bufio.NewScanner(measurementsFile)

	m := make(map[string]Measurement)
	for scanner.Scan() {
		text := scanner.Text()
		readings := strings.Split(text, ";")
		temp, err := strconv.ParseFloat(readings[1], 64)
		if err != nil {
			panic(err)
		}
		
		lowest := m[readings[0]].Min
		highest := m[readings[0]].Min

		if m[readings[0]].Count == 0 {
			lowest = 100
			highest = -100
		}

		if lowest > temp {
			lowest = temp
		}

		if highest < temp {
			highest = temp
		}

		m[readings[0]] = Measurement{
			Count: m[readings[0]].Count + 1,
			Sum:   m[readings[0]].Sum + temp,
			Max:   highest,
			Min:   lowest,
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for k,v := range m {
		fmt.Printf("%s=%.1f/%.1f/%.1f\n", k, v.Min, v.Sum/float64(v.Count), v.Max)
	}

    fmt.Println("Run time:", time.Since(start))
}
