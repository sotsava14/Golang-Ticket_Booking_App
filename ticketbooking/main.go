package main

import "fmt"

func main() {
	var n int
	var arr [10]string
	var ans string
	var mov string
	t := 50
	fmt.Println("Welcome to BooMyShow!")
	fmt.Print("Enter the no.of tickets:")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Person", i+1, " Name:")
		fmt.Scan(&arr[i])
	}
	isInValid := n > 50 || n <= 0
	if isInValid {
		fmt.Println("Tickets not available")
		return
	}
	t = t - n
	if t <= 10 {
		fmt.Println("Your Tickets have been booked! (Fast-filling)Tickets available:", t)
	} else {
		fmt.Println("Your Tickets have been booked!Tickets available:", t)
	}
	fmt.Println("Do you need Popcorn?(Yes or No)")
	fmt.Scan(&ans)
	switch ans {
	case "Yes":
		fmt.Println("Visit our Homepage for the best popcorn!")
	case "No":
		fmt.Println("Enjoy the Movie:)")
	}
	fmt.Println("Movie Name:")
	fmt.Scan(&mov)
	bye(mov)
}

func bye(m string) {
	fmt.Println("THANK YOU! ENJOY WATCHING ", m)
}
