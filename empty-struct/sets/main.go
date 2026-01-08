package main

import "fmt"

func main() {
	fruits := map[string]struct{}{
		"apple": {},
		"banana": {},
		"grapes": {},
		"papaya": {},
	}

	fruit := "apple"
	_, exists := fruits[fruit]
	if exists {
		fmt.Printf("%s fruit exists in set", fruit)
	}


	fmt.Println("listing fruits")

	for f := range fruits{
		fmt.Println(f)
	}

	fruits["kiwi"] = struct{}{}

	fmt.Println("added a new fruit, kiwi")

	fmt.Println("listing fruits")

	for f := range fruits{
		fmt.Println(f)
	}
}