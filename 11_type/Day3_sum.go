package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40, 50))
	fmt.Println(sum(10, 20, "30", 40, 50))
	fmt.Println(sum(10, 20, "30", 40, 50, "abc"))
	fmt.Println(sum(10, 20, []int{30, 40, 50}))
	fmt.Println(sum(10, 20, []interface{}{"30", 40, 50, "abc"}))
}

func sum(args ...interface{}) int {
	result := 0
	for _, arg := range args {
		switch value := arg.(type) {
		case int:
			result += value
		case string:
			if intVal, err := strconv.Atoi(value); err == nil {
				result += intVal
			}
		case []int:
			data := make([]interface{}, len(value))
			for idx, v := range value {
				data[idx] = v
			}
			result += sum(data...)
		case []interface{}:
			result += sum(arg.([]interface{})...)
		}
	}
	return result
}
