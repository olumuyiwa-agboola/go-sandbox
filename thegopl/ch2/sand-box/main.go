package main

import "fmt"

// var p = f()

// func f() *int {
// 	v := 1
// 	return &v
// }

// func incr(p *int) int {
// 	*p++
// 	return *p
// }

func main() {
	// s := ""
	// fmt.Println(s)

	// x := 1
	// p := &x
	// fmt.Println(p)
	// fmt.Println(*p)
	// *p = 2
	// fmt.Println(x)

	// var x, y int
	// fmt.Println(&x == &x, &x == &y, &x == nil)

	// fmt.Println(f() == f())

	// v := 1
	// incr(&v)
	// fmt.Println(v)
	// fmt.Println(incr(&v))

	p := new(int)
	fmt.Printf("p is %v\n", p)
	fmt.Printf("*p is %v\n", *p)

	*p = 2
	fmt.Printf("After executing *p = 2, p is %v\n", p)
	fmt.Printf("After executing *p = 2, *p is %v\n", *p)
}
