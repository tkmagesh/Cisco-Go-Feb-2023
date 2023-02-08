package main

import "fmt"

func main() {
	//if else
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	//no variable scope is limited to the if-else block
	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	/* switch case */
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 	=> "Excellent"
		otherwise 	=> "Invalid Score"
	*/

	/* score := 6
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
		fmt.Println("Invalid Score")
	}
	*/

	switch score := 6; score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid Score")
	}

	//switch case (if else if)
	switch no := 21; {
	case no%2 == 0:
		fmt.Printf("%d is an even number\n", no)
	case no%2 != 0:
		fmt.Printf("%d is an odd number\n", no)
	}

	//fallthrough
	fmt.Println("Switch-Case with fallthrough")
	switch n := 3; n {
	case 0:
		fmt.Println("n = 0")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		// fallthrough
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
		fallthrough
	case 9:
		fmt.Println("n <= 9")
		fallthrough
	case 10:
		fmt.Println("n <= 10")
	}

	fmt.Println("Switch Case with fallthrough (applied)")
	switch plan := "Premium"; plan {
	case "Super":
		fmt.Println("[Super] - for a family of 4")
		fallthrough
	case "Premium":
		fmt.Println("[Premium] - Download")
		fmt.Println("[Premium] - Playlist")
		fallthrough
	case "Pro":
		fmt.Println("[Pro] - No ads!")
		fallthrough
	case "Free":
		fmt.Println("[Free] - Listen to songs")
	}
}
