package main

import (
	"fmt"
)

// custom print function to print line with single string
func print(s string) {
	fmt.Println(s)
}

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

// // //if else logic flow/branching
// func main() {
// 	// basic if else flow
// 	if 7%2 == 0 {
// 		fmt.Printf("7 is even")
// 	} else {
// 		fmt.Printf("7 is odd")
// 	}
// 	//if no else, single conditional
// 	if 23*2 > 12 {
// 		fmt.Println("\ntrue facts")
// 	}
// 	//boolean comparisions within conditions
// 	if 23%2 == 0 || false && true {
// 		fmt.Println("uhhhhh not sure")
// 	} else {
// 		fmt.Println("this is the else :)")
// 	}

// 	// statement can precede conditional?
// 	if num := 11; num < 12 {
// 		fmt.Println("less 12")
// 	} else {
// 		fmt.Println("greater than or equal to 12")
// 	}

// }

// // basic switch statements
// func main() {

// 	//simple switch evaluating on one condition
// 	i := 0
// 	fmt.Print("Write ", i, " as ")
// 	switch i {
// 	case 0:
// 		print("zero")
// 	case 1:
// 		print("one")
// 	case 2:
// 		print("two")
// 	}
// 	//more complex switch, multiple conditions for one output
// 	switch time.Now().Weekday() {
// 	case time.Sunday, time.Saturday:
// 		print("weekend")
// 	default:
// 		print("not weekend")
// 	}

// 	//alternative if statement. Blank switch creates a conditonial logic flow within brackets
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		print("erly burd gits wurm")
// 	default:
// 		print("cap")
// 	}

// 	//takes interface and returns its type
// 	// function paired with switch to compare datatype attributes, interface.() 'type' in this case
// 	typeFunc := func(i interface{}) {
// 		switch t := i.(type) {
// 		case int:
// 			print("int")
// 		case bool:
// 			print("i am bool")
// 		default:
// 			fmt.Printf("Dont know type %T\n", t)
// 		}
// 	}

// 	typeFunc(12)
// 	typeFunc(true)
// 	typeFunc("stringy")
// }
