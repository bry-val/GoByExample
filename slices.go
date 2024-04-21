package main

import "fmt"

//practice with slices.
//interface for sequences, potentially more powerful than arrays
func main() {
	//slices are able to be unitilizaed and have no predefined size attribute . The only attribute is the type of its members
	var s []string
	fmt.Println("uninitilzaied slice:", s, "\nlength of slice:", len(s), "\nis nil?:", s == nil)

	//empty slice initilizaed to predetermined size, make([]arrayType, lenght)
	//capacity of slice method => cap()
	s = make([]string, 3)
	fmt.Println("initiialized slice:", s, "\nlength of slice:", len(s), "\nis nil?:", s == nil, "\ncapacity of slice?", cap(s))

}
