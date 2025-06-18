/* We can use range for iterating over a range of indexes in arrays, slices, or maps. Range provides keys and
values in variables  */

package main

import "fmt"

func main() {
	var slice1 = []int{9, 8, 7}

	for key, value := range slice1 {
		fmt.Println(key, value)
	}

	var map1 = map[string]string{"Name": "John Doe"}
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	var string1 = "John Doe"
	for k, v := range string1 {
		fmt.Println(k, v)
	}
}
