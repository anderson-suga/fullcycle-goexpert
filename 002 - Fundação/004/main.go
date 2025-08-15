package main

import "fmt"

const a = "Hello World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Anderson"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo da variável do E é %T\n", e)
	fmt.Printf("O tipo da variável do F é %T\n", f)
}
