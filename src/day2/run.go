package main

import (
    "bufio"
    "fmt"
    "log"
	"strings"
	"strconv"
    "os"
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

func get_color_counts(input string) bool {
	//Game 28: 5 blue, 6 green, 1 red; 13 blue; 1 red, 9 blue, 10 green
	fmt.Println(input)
	colors := strings.Split(input, ":")[1]
	parts := strings.Split(colors, ";")
	for _, part := range parts {
		r, g, b := 0, 0, 0
		color := strings.Split(part, ",")
		for _, c := range color {
			stat := strings.Split(c, " ")
			c_count, _ := strconv.Atoi(stat[1])
			if stat[2] == "red" {
				r += c_count
			} else if stat[2] == "blue" {
				b += c_count
			} else if stat[2] == "green" {
				g += c_count
			}
		}
		if r > 12 || g > 13 || b > 14 {
			fmt.Println("Yes: ", "r:", r, " g: ", g, " b: ", b)
			return false
		}
	}
	return true
}

func get_max_color_multiply(input string) int {
	//Game 28: 5 blue, 6 green, 1 red; 13 blue; 1 red, 9 blue, 10 green
	fmt.Println(input)
	r, g, b := 0, 0, 0
	colors := strings.Split(input, ":")[1]
	parts := strings.Split(colors, ";")
	for _, part := range parts {
		color := strings.Split(part, ",")
		for _, c := range color {
			stat := strings.Split(c, " ")
			c_count, _ := strconv.Atoi(stat[1])
			if stat[2] == "red" && c_count > r {
				r = c_count
			} else if stat[2] == "blue" && c_count > b {
				b = c_count
			} else if stat[2] == "green" && c_count > g {
				g = c_count
			}
		}
	}
	return r * g * b
}

func run_sec_1() int {
	allText := get_file_content_as_array("data_2.txt")
	total := 0
	for i, l := range allText {
		game := i + 1
		valid := get_color_counts(l)
		//12 red cubes, 13 green cubes, and 14 blue cubes?
		if valid {
			total += game
		}
	}
	return total
}

func run_sec_2() int {
	allText := get_file_content_as_array("data_2.txt")
	total := 0
	for _, l := range allText {
		mult := get_max_color_multiply(l)
		total += mult
	}
	return total
}

func main() {
	ans1 := run_sec_1()
	fmt.Println("ans1: ", ans1)
	ans2 := run_sec_2()
	fmt.Println("ans2: ", ans2)
}