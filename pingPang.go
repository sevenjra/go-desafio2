package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0 :
			fmt.Println("PingPang")
		case 	i%3 == 0:
			fmt.Println("Ping")
		case 	i%5 == 0:
			fmt.Println("Pang")
		default:
			fmt.Println(i)	
		}
		
	}
}