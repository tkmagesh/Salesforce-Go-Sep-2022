package main

import "fmt"

func main() {
	/* if else */
	fmt.Printf("if else\n")
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	/* switch case */
	fmt.Printf("\nswitch case\n")
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/
	score := 6
	/*
		switch score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	/* if elsif alternative */
	/*
		no := 21
		switch {
		case no%2 == 0:
			fmt.Printf("%d is an even number\n", no)
		case no%2 != 0:
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	switch no := 21; {
	case no%2 == 0:
		fmt.Printf("%d is an even number\n", no)
	case no%2 != 0:
		fmt.Printf("%d is an odd number\n", no)
	}

	fmt.Println("Switch case with fallthrough")

	switch n := 4; n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
	}

	switch plan := "super"; plan {
	case "super":
		fmt.Println("All super features")
		fallthrough
	case "premium":
		fmt.Println("All premium features")
		fallthrough
	case "pro":
		fmt.Println("All pro features")
		fallthrough
	case "free":
		fmt.Println("All free features")
	}

	/* for */
	fmt.Printf("\nfor\n")
	fmt.Println("v1.0")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\nwhile (v2.0)\n")
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %d\n", numSum)

	fmt.Printf("\ninfinite (v3.0)\n")

	num := 1
	for {
		num += num
		if num > 100 {
			break
		}
	}
	fmt.Printf("num = %d\n", num)

	fmt.Println("Using labels")

OUTER_LOOP:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				fmt.Println("========================")
				continue OUTER_LOOP
			}
		}
	}
}
