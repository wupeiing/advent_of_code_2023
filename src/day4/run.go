package main

import (
    "bufio"
    "fmt"
    "log"
	"strings"
	"strconv"
    "os"
	"math"
)

func get_file_content_as_array(filename string) []string {
	arr := make([]string, 0)
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
		arr = append(arr, scanner.Text())
    }
	return arr
}

func getWinningNumbers(winNumStr string) map[int]int {
	winNums := make(map[int]int, 0)
	sections := strings.Split(winNumStr, " ")
	for _, s := range sections {
		v, e := strconv.Atoi(s)
		if e == nil {
			winNums[v] += 1
		}
	}
	return winNums
}

func getPoints(winNums map[int]int, nums string, isPoint bool) int {
	match := 0
	sections := strings.Split(nums, " ")
	for _, s := range sections {
		v, e := strconv.Atoi(s)
		if e == nil {
			if winNums[v] != 0 {
				match += 1
			}
		}
	}
	if !isPoint {
		return match
	}
	if match > 0 {
		target := float64(match - 1)
		return int(math.Pow(2, target))
	}
	return 0
}

func run_sec_1() int {
	allText := get_file_content_as_array("data_4.txt")
	total := 0
	for _, l := range allText {
		parts := strings.Split(l, "|")
		winNums := getWinningNumbers(parts[0])
		points := getPoints(winNums, parts[1], true)
		total += points
	}
	return total
}

func run_sec_2() int {
	allText := get_file_content_as_array("data_4.txt")
	n := len(allText)
	arr := make([]int, n)
	total := 0
	for i, l := range allText {
		arr[i] += 1
		parts := strings.Split(l, "|")
		winNums := getWinningNumbers(parts[0])
		wins := getPoints(winNums, parts[1], false)
		for j := 1; j<=wins; j++ {
			if i + j >= n {
				break
			}
			arr[i+j] += arr[i]
		}
		total += arr[i]
	}
	return total
}

// c1 := 1
// c2 :=

func main() {
	ans1 := run_sec_1()
	fmt.Println("ans1: ", ans1)
	ans2 := run_sec_2()
	fmt.Println("ans2: ", ans2)
}
