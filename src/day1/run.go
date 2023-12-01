package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func get_digit_when_valid(i int, m string) (int, bool) {
	n := len(m)
	if m[i] >= '0' && m[i] <= '9' {
		return int(m[i] - '0'), true
	} else if m[i] == 'o' {
		if i + 3 <= n && string(m[i:i+3]) == "one" {
			return 1, true
		}
	} else if m[i] == 't' {
		if i + 3 <= n && string(m[i:i+3]) == "two" {
			return 2, true
		} else if i + 5 <= n && string(m[i:i+5]) == "three" {
			return 3, true
		}
	} else if m[i] == 'f' {
		if i + 4 <= n && string(m[i:i+4]) == "four" {
			return 4, true
		} else if i + 4 <= n && string(m[i:i+4]) == "five" {
			return 5, true
		}
	} else if m[i] == 's' {
		if i + 3 <= n && string(m[i:i+3]) == "six" {
			return 6, true
		} else if i + 5 <= n && string(m[i:i+5]) == "seven" {
			return 7, true
		}
	} else if m[i] == 'e' {
		if i + 5 <= n && string(m[i:i+5]) == "eight" {
			return 8, true
		}
	} else if m[i] == 'n' {
		if i + 4 <= n && string(m[i:i+4]) == "nine" {
			return 9, true
		}
	}
	return 0, false
}

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

func run_sec_1() int {
	allText := get_file_content_as_array("data_1.txt")
	total := 0
	for _, n := range allText {
		a, b := -1, -1
		for i, _ := range n {
			if n[i] >= '0' && n[i] <= '9' {
				if a == -1 {
					a = int(n[i] - '0')
				}
				b = int(n[i] - '0')
			}
		}
		total += a * 10 + b
	}
	return total
}

func run_sec_2() int {
	allText := get_file_content_as_array("data_1.txt")
	total := 0
	for _, m := range allText {
		a, b := -1, -1
		for i, _ := range m {
			d, valid := get_digit_when_valid(i, m)
			if valid {
				if a == -1 {
					a = d
				}
				b = d
			}
		}
		total += a * 10 + b
	}
	return total
}

func main() {
	ans1 := run_sec_1()
	fmt.Println("ans1: ", ans1)
	ans2 := run_sec_2()
	fmt.Println("ans2: ", ans2)
}