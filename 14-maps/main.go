/* Maps in Go are similar to Dictionaries, which store key-value pairs. We can define the type of keys as well as values
when creating a map. If a key does not exist in a map, it returns a zero value. */

package main

import "fmt"

func main() {
	var map1 = make(map[int]string)
	map1[1] = "Pokemon"
	map1[2] = "Digimon"
	fmt.Println(map1)
	delete(map1, 2)
	fmt.Println(map1)
	clear(map1)
	fmt.Println(map1)

	value, okay := map1[1]
	if okay {
		fmt.Println("Exists", value)
	} else {
		fmt.Println("Doesn't")
	}
}
