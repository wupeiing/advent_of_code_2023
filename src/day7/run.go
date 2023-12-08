package main

import (
    "bufio"
    "fmt"
    "log"
	"strings"
	"strconv"
    "os"
	"sort"
)

func get_file_content_as_array(filename string) []string {
	arr := make([]string, 0)
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		arr = append(arr, scanner.Text())
    }
	return arr
}

func getCardResult(c string) int {
	mp := make(map[byte]int, 0)
	max := 0
	for i, _ := range c {
		mp[c[i]] += 1
		if mp[c[i]] > max {
			max = mp[c[i]]
		}
	}
	if max == 5 {
		return 6
	} else if max == 4 {
		return 5
	} else if max == 3 {
		if len(mp) == 2 {
			return 4
		} else {
			return 3
		}
	} else if max == 2 {
		if len(mp) == 3 {
			return 2
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func getCardResult2(c string) int {
	mp := make(map[byte]int, 0)
	max := 0
	j_count := 0
	for i, _ := range c {
		if c[i] == 'J' {
			j_count += 1
			continue
		}
		mp[c[i]] += 1
		if mp[c[i]] > max {
			max = mp[c[i]]
		}
	}
	if j_count != 0 {
		max += j_count
	}
	if max == 5 {
		return 6
	} else if max == 4 {
		return 5
	} else if max == 3 {
		if len(mp) == 2 {
			return 4
		} else {
			return 3
		}
	} else if max == 2 {
		if len(mp) == 3 {
			return 2
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func run_sec_2() int {
	allText := get_file_content_as_array("data_7.txt")
	total := 0
	mp := make(map[string]int, 0)
	arr := make([][]string, 0)
	for i := 0; i<7; i++ {
		tmp := make([]string, 0)
		arr = append(arr, tmp)
	}
	for _, l := range allText {
		parts := strings.Split(l, " ")
		num, _ := strconv.Atoi(parts[1])
		mp[parts[0]] = num
		res := getCardResult2(parts[0])
		arr[res] = append(arr[res], parts[0])
	}
	multiply := 1
	orderList := "AKQT98765432J"
	for _, g := range arr {
		if len(g) == 0 {
			continue
		}
		sort.Slice(g, func(i, j int) bool {
			// Compare each character separately
			for k := 0; k < len(g[i]) && k < len(g[j]); k++ {
				posI := strings.Index(orderList, string(g[i][k]))
				posJ := strings.Index(orderList, string(g[j][k]))

				// If positions are different, return the result
				if posI != posJ {
					return posI > posJ
				}
			}

			// If all characters are the same up to the minimum length, consider length
			return len(g[i]) < len(g[j])
		})
		for _, s := range g {
			total += mp[s] * multiply
			multiply += 1
		}
	}
	return total
}

func run_sec_1() int {
	allText := get_file_content_as_array("data_7.txt")
	total := 0
	mp := make(map[string]int, 0)
	arr := make([][]string, 0)
	for i := 0; i<7; i++ {
		tmp := make([]string, 0)
		arr = append(arr, tmp)
	}
	for _, l := range allText {
		parts := strings.Split(l, " ")
		num, _ := strconv.Atoi(parts[1])
		mp[parts[0]] = num
		res := getCardResult(parts[0])
		arr[res] = append(arr[res], parts[0])
	}
	multiply := 1
	orderList := "AKQJT98765432"
	for _, g := range arr {
		if len(g) == 0 {
			continue
		}
		sort.Slice(g, func(i, j int) bool {
			// Compare each character separately
			for k := 0; k < len(g[i]) && k < len(g[j]); k++ {
				posI := strings.Index(orderList, string(g[i][k]))
				posJ := strings.Index(orderList, string(g[j][k]))

				// If positions are different, return the result
				if posI != posJ {
					return posI > posJ
				}
			}

			// If all characters are the same up to the minimum length, consider length
			return len(g[i]) < len(g[j])
		})
		for _, s := range g {
			total += mp[s] * multiply
			multiply += 1
		}
	}
	return total
}

func main() {
	ans1 := run_sec_1()
	fmt.Println("ans1: ", ans1)
	ans2 := run_sec_2()
	fmt.Println("ans2: ", ans2)
}
