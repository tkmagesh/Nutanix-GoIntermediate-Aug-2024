package main

import "fmt"

func main() {
	increment := GetIncrementer()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
	fmt.Println(increment()) //=> 5
}

func GetIncrementer() func() int { // step 1
	var count int             // step 2
	increment := func() int { //step - 3
		count++ // step - 4
		return count
	}
	return increment // step - 5
}
