package main

import (
	"fmt"
	"strings"
)

func main() {
	var ba [][]string
	for {
		var row string
		if _, err := fmt.Scanln(&row); err != nil {
			break
		}
		ba = append(ba, strings.Split(row, ""))
	}
	tr := 0
	rc := 0
	r1 := 0
	for {
		rtr := 0
		rc++
		result := make([][]string, len(ba))
		for i := 0; i < len(ba); i++ {
			result[i] = make([]string, len(ba[i]))
			for j := 0; j < len(ba[i]); j++ {
				dirs := []int{-1, 0, 1}
				count := 0
				for _, dy := range dirs {
					for _, dx := range dirs {
						if dy == 0 && dx == 0 {
							continue
						}
						nr := i + dy
						nc := j + dx
						if nr >= 0 && nr < len(ba) && nc >= 0 && nc < len(ba[i]) {
							if ba[nr][nc] == "@" {
								count += 1
							}
						}
					}
				}
				if ba[i][j] == "@" && count < 4 {
					result[i][j] = "x"
					rtr++
				} else {
					result[i][j] = ba[i][j]
				}
			}
		}
		tr += rtr
		if rc == 1 {
			r1 += rtr
		}
		if rtr == 0 {
			break
		}
		ba = result
	}
	fmt.Println(r1, tr)
}
