package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	low, high int
}

func isFresh(id int, ranges [][]string) bool {
	for x := range ranges {
		low, _ := strconv.Atoi(ranges[x][0])
		high, _ := strconv.Atoi(ranges[x][1])
		if id >= low && id <= high {
			return true
		}
	}
	return false
}

func main() {
	var ingr [][]string
	var ba []string

	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, "-") {
			ingr = append(ingr, strings.Split(row, "-"))
		} else if row != "" {
			ba = append(ba, row)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for x := range ba {
		id, _ := strconv.Atoi(ba[x])
		if isFresh(id, ingr) {
			count++
		}
	}

	ranges := make([]Range, len(ingr))
	for i := range ingr {
		low, _ := strconv.Atoi(ingr[i][0])
		high, _ := strconv.Atoi(ingr[i][1])
		ranges[i] = Range{low, high}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].low < ranges[j].low
	})

	merged := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		current := ranges[i]

		if current.low <= last.high+1 {
			if current.high > last.high {
				last.high = current.high
			}
		} else {
			merged = append(merged, current)
		}
	}

	totalFresh := 0
	for _, r := range merged {
		totalFresh += r.high - r.low + 1
	}

	fmt.Println(count)
	fmt.Println(totalFresh)
}
