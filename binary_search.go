package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func get_value_from_line(line_number int, line_length int, input_file *os.File) int64 {
	current_line := make([]byte, line_length)

	input_file.Seek(int64((line_length+1)*line_number), 0)
	input_file.Read(current_line)
	string_values := strings.Split(string(current_line), ";")
	int_value, _ := strconv.Atoi(string_values[0])

	return int64(int_value)
}

func main() {
	input_file, _ := os.Open("test_data.csv")

	scanner := bufio.NewScanner(input_file)
	scanner.Scan()
	line_length := utf8.RuneCountInString(scanner.Text())
	num_lines, _ := strconv.Atoi(scanner.Text())
	high_end := 1
	low_end := num_lines
	current_index := 0
	var requested_value int64 = 4

	done := false
	for !done {
		if (low_end - high_end) <= 1 {
			current_index = -1
			done = true
		} else {
			current_index = (high_end + low_end) / 2
			int_value := get_value_from_line(current_index, line_length, input_file)
			if int_value == requested_value {
				done = true
			} else {
				if int_value < requested_value {
					high_end = current_index
				} else {
					low_end = current_index
				}
			}
		}
	}

	if current_index > 0 {
		done = false
		for !done {
			current_index--
			int_value := get_value_from_line(current_index, line_length, input_file)
			if int_value != requested_value {
				done = true
			}
		}
		fmt.Printf("%d\n", current_index+1)
	} else {
		fmt.Printf("Not found!\n")
	}
}
