package IotaDemo

import (
	"fmt"
)

const(
	a = iota
	b
	c
)

const(
	m=1<<iota
	n=2<<iota
	x=10
	y = iota
	z=iota>>1
	o
)

func IotaTest() {

	fmt.Println("a=",a)
	fmt.Println("b=",b)
	fmt.Println("c=",c)

	fmt.Println("m=",m)
	fmt.Println("n=",n)
	fmt.Println("x=",x)
	fmt.Println("y=",y)
	fmt.Println("z=",z)
	fmt.Println("o=",o)
}