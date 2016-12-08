package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Parameter: %s", strings.Split(r.URL.Path, "/")[2])
	parameter, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	parameter64 := int64(parameter)
	ret := find_in_sorted_file("test_data.csv", parameter64)
	fmt.Printf("%s\n", ret)
	fmt.Fprintf(w, "Result: %s", ret)
}

func get_values_from_line(line_number int, line_length int, input_file *os.File) (int64, int64) {
	current_line := make([]byte, line_length)

	input_file.Seek(int64((line_length+1)*line_number), 0)
	input_file.Read(current_line)
	string_values := strings.Split(string(current_line), ";")
	int_value, _ := strconv.Atoi(string_values[0])
	related_value, _ := strconv.Atoi(string_values[1])

	return int64(int_value), int64(related_value)
}

func find_in_sorted_file(file_name string, search_num int64) string {
	input_file, _ := os.Open(file_name)

	scanner := bufio.NewScanner(input_file)
	scanner.Scan()
	line_length := utf8.RuneCountInString(scanner.Text())
	num_lines, _ := strconv.Atoi(scanner.Text())
	high_end := 1
	low_end := num_lines
	current_index := 0

	done := false
	for !done {
		if (low_end - high_end) <= 1 {
			current_index = -1
			done = true
		} else {
			current_index = (high_end + low_end) / 2
			int_value, _ := get_values_from_line(current_index, line_length, input_file)
			if int_value == search_num {
				done = true
			} else {
				if int_value < search_num {
					high_end = current_index
				} else {
					low_end = current_index
				}
			}
		}
	}

	fmt.Printf("%d\n", current_index)
	var ret_strings []string
	if current_index > 0 {
		done = false
		for !done {
			current_index--
			int_value, _ := get_values_from_line(current_index, line_length, input_file)
			if int_value != search_num {
				done = true
			}
		}
		current_index += 1
		done = false
		for !done {
			int_value, related_value := get_values_from_line(current_index, line_length, input_file)
			fmt.Printf("%d - %d\n", int_value, related_value)
			if int_value != search_num {
				done = true
			} else {
				str_value := fmt.Sprintf("%d", related_value)
				ret_strings = append(ret_strings, str_value)
				current_index += 1
			}
		}
	}

	ret_value := strings.Join(ret_strings, "\n")
	return ret_value
}

func main() {
	http.HandleFunc("/related/", handler)
	http.ListenAndServe(":8080", nil)
}
