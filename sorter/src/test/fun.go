package main

import "fmt"

func add(a int,b int)(c int)  {
	c = a+b
	return
}
func main()  {
	a,b := 1,2
	c := add(a,b)
	fmt.Println(c)
}
