package main

import "fmt"

func main() {
	fmt.Print("Mesa")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println(" linha")

	x := 3.141516

	// 	fmt.Print("O valor de X é " + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de X é " + xs)
	fmt.Println("O valor de X é", x)

	fmt.Printf("O valor de X é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n %d %f %.1f %v %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v ", a, b, c, d)

}
