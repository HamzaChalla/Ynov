package main

import "fmt"

func main() {
	var s string =********

	fmt.Println("Saisissez un mot")
	fmt.Scanln(&s)
	
	fmt.Println(string(s[:8]))
	fmt.Println(string(s[:7]))
	fmt.Println(string(s[:6]))
	fmt.Println(string(s[:5]))
	fmt.Println(string(s[:4]))
	fmt.Println(string(s[:3]))
	fmt.Println(string(s[:2]))
	fmt.Println(string(s[:1]))
	fmt.Println(string(s[0]))

	}

}
