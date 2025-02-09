package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a:=0
	b:=1
	c:=1
	
	return func() int{
		x := b+c;
		a=b;
		b=c;
		c=x;
		return a;
	}
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

