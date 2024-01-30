package main

import "fmt"

// hello world

// func main() {
// 	fmt.Println("hello world")
// }

// values
// func main() {
// 	fmt.Println("oy" + "momma")
// 	fmt.Println("oy", 12)

// 	fmt.Println(true && false)
// 	fmt.Println(true || false)
// 	fmt.Println(!false)
// }

//variables

// func main() {
// 	nWords := "stringify"

// 	var five = 5

// 	fmt.Println(nWords)
// 	fmt.Println(five)

// }

// //constants
// func main() {
// 	const s string = "tumadre"

// 	val := 84

// 	const num = 5461154156

// 	fmt.Println(num + val)
// }

// // // for loop, only looping construct (no while loop)
// func main() {
// 	// basic for loop, single condition
// 	i := 0
// 	for i <= 3 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}

// 	//classic initializer, condition, increment loop
// 	for j := 0; j <= 3; j++ {
// 		fmt.Println(j)
// 	}

// 	// inifintie loop (while true equivalent)
// 	// must be broken by break statement
// 	for {
// 		fmt.Println("Looooop")
// 		break
// 	}

// 	//classic loop with inner condition using continue keyword to skip to next iteration
// 	for n := 0; n <= 5; n++ {
// 		if n%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(n)
// 	}
// }

// //if else logic flow/branching
func main() {
	// basic if else flow
	if 7%2 == 0 {
		fmt.Printf("7 is even")
	} else {
		fmt.Printf("7 is odd")
	}
	//if no else, single conditional
	if 23*2 > 12 {
		fmt.Println("\ntrue facts")
	}
	//boolean comparisions within conditions
	if 23%2 == 0 || false && true {
		fmt.Println("uhhhhh not sure")
	} else {
		fmt.Println("this is the else :)")
	}

}
