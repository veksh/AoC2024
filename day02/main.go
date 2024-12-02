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
	return(len(reports))
}

func main() {
	reports := readInput()
	fmt.Println(reports)
	fmt.Println("ans1:", ans1(reports))
}