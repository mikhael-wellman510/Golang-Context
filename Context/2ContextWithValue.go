package main

import (
	"context"
	"fmt"
)

func main() {

	contextA := context.Background() // Parent

	contextB := context.WithValue(contextA, "b", "B") // Sub Parent A
	contextC := context.WithValue(contextA, "c", "C") // Sub Parent A

	contextD := context.WithValue(contextB, "d", "D") //  Child B
	contextE := context.WithValue(contextB, "e", "E") // child B

	contextDD := context.WithValue(contextD, "dd", "DD") // Child child D
	contextEE := context.WithValue(contextD, "ee", "EE") // Child Child D

	contextF := context.WithValue(contextC, "f", "F") // Child C

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextDD)
	fmt.Println(contextEE)

	fmt.Println(contextF.Value("f"))  // Dapat
	fmt.Println(contextEE.Value("d")) // Dapat karena ada di Parent nya
	fmt.Println(contextF.Value("b"))  // Tidak Dapat - Beda Parent
	fmt.Println(contextB.Value("e"))  // Tidak Dapat Mengambil Child
}
