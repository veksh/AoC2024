package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)


func readInput() [][]int {
	fName := "input_test.txt"
	if len(os.Args) >= 2 {
		fName = os.Args[1]
	}
	fh, err := os.Open(fName); if err != nil {log.Fatal("cannot open file", err)}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	res := [][]int{}
	for scanner.Scan() {
		strnums := strings.Split(scanner.Text(), " ")
		nums := lo.Map(strnums, func(s string, _ int) int {n, _ := strconv.Atoi(s); return n})
		res = append(res, nums)
	}
	return res
}


func ans1(reports [][]int) int {
	res := 0
	for _, rep := range(reports) {
		sign := 1
		if rep[1] < rep[0] {
				sign = -1
		}
		ok := 1
		prevlevel := rep[0] - 1*sign
		for _, level := range(rep) {
			diff := (level - prevlevel)/sign
			if diff >= 1 && diff <= 3 {
				prevlevel = level
			} else {
				ok = 0
				break
			}
		}
		res += ok
	}
	return(res)
}

func ans2(reports [][]int) int {
	res := 0
	for _, rep := range(reports) {
		ok, bad := 1, 0
		diff0 := max(rep[0], rep[1]) - min(rep[0], rep[1])
		if diff0 < 1 || diff0 > 3 {
			bad = 1
			rep = rep[1:]
		}
		sign := 1
		if rep[1] < rep[0] {
				sign = -1
		}
		prevlevel := rep[0] - 1*sign
		for _, level := range(rep) {
			diff := (level - prevlevel)/sign
			if diff >= 1 && diff <= 3 {
				prevlevel = level
			} else {
				if bad == 0 {
					bad = 1
				} else {
					ok = 0
					break
				}
			}
		}
		res += ok
	}
	return(res)
}

func main() {
	reports := readInput()
	fmt.Println("ans1:", ans1(reports))
	fmt.Println("ans2:", ans2(reports))
}