package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strings"
	"strconv"
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
        //fmt.Println(scanner.Text())
		arr = append(arr, scanner.Text())
    }
	return arr
}

func getNumArray(line string) []int {
	arr := make([]int, 0)
	ns := strings.Split(line, " ")
	for _, n := range ns {
		v, _ := strconv.Atoi(n)
		arr = append(arr, v)
	}
	return arr
}

func needMoreLoop(arr []int) bool {
	mp := make(map[int]int, 0)
	for _, x := range arr {
		mp[x] += 1
	}
	return len(mp) == 1
}

func getNextLine(arr []int) []int {
	tmp := make([]int, 0)
	for i:=1;i<len(arr);i++ {
		tmp = append(tmp, arr[i] - arr[i-1])
	}
	return tmp
}

func getNextNum(arr []int, backward bool) int {
	arrs := make([][]int, 0)
	arrs = append(arrs, arr)
	run := needMoreLoop(arr)
	for !run {
		next := getNextLine(arrs[len(arrs) - 1])
		arrs = append(arrs, next)
		run = needMoreLoop(next)
	}
	if backward {
		for j:=len(arrs)-1;j>0; j-- {
			app := arrs[j][len(arrs[j])-1]
			arrs[j-1] = append(arrs[j-1], arrs[j-1][len(arrs[j-1])-1] + app)
		}
		fmt.Println(arrs)
		return arrs[0][len(arrs[0])-1]
	} else {
		for j:=len(arrs)-1;j>0; j-- {
			app := arrs[j][0]
			arrs[j-1] = append([]int{arrs[j-1][0] - app}, arrs[j-1]...)
		}
		fmt.Println(arrs)
		return arrs[0][0]
	}

}

func run_sec_1() int {
	allText := get_file_content_as_array("data_9.txt")
	total := 0
	for _, l := range allText {
		arr := getNumArray(l)
		// fmt.Println(arr)
		nextNum := getNextNum(arr, true)
		total += nextNum
	}
	return total
}

func run_sec_2() int {
	allText := get_file_content_as_array("data_9.txt")
	total := 0
	for _, l := range allText {
		arr := getNumArray(l)
		// fmt.Println(arr)
		nextNum := getNextNum(arr, false)
		total += nextNum
	}
	return total
}

func main() {
	ans1 := run_sec_1()
	fmt.Println("ans1: ", ans1)
	ans2 := run_sec_2()
	fmt.Println("ans2: ", ans2)
}