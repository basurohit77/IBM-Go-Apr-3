package main

import (
	"fmt"
	"strconv"
)


var sum = 0

func main() {
	var x interface{}
	//x = true
	/* x = "abc"
	x = true
	fmt.Println(x) */
	x = 100
	//x = "abc"
	fmt.Printf("Type of x = %T\n", x)
	if value, ok := x.(int); ok {
		fmt.Println(value + 100)
	} else {
		fmt.Printf("%v is not int type\n", x)
	}

	switch v := x.(type) {
	case int:
		fmt.Printf("Twice of %v is %v\n", v, 2*v)
	case string:
		fmt.Printf("Len of %v is %d\n", v, len(v))
	default:
		fmt.Println("Unknown type of %v = %T\n", v, v)
	}

	fmt.Println(add(10, 20, 30, 40, 50, "abc"))

	fmt.Println("add(10, 20, '30', 40, 50" )
	fmt.Println(add(10, 20, "30", 40, 50))

	fmt.Println("add(10, 20, []int{30,40,50}" )
	fmt.Println(add(10,20,[]int{30,40,50}))

	fmt.Println("sum(10,20,[]interface{}{'30',40,50,abc}")
	fmt.Println(add(10,20,[]interface{}{"30",40,50,"abc"}))
}


func add(nos ...interface{}) int {

    var y interface{}

	for i := 0; i < len(nos); i++ {
		y = nos[i]

		switch v := y.(type) {
		case int:
			sum += v
		case string:
			k, err := strconv.Atoi(v)
			if err == nil {
				sum += k
			}
		case []int:
			vlen := len(v)
			for j := 0; j < vlen; j++ {
				sum += v[j]
			}
		case []interface{}:
			add(v...)
		default:
			fmt.Println("Unknown type of %v = %T\n", v, v)
		}

	}
	return sum

}
/*
sum() => 0
sum(10) => 10
sum(10,20) => 30
sum(10,20,30,40,50) => 150
sum(10,20,"30",40,50) => 150
sum(10,20,"30",40,50,"abc") => 150
sum(10,20,[]int{30,40,50}) => 150
sum(10,20,[]interface{}{"30",40,50,"abc"}) => 150
*/


