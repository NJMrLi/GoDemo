package StringDemo

import (
	"fmt"
)


func StringTest1() {
	str := "Hello,World"
	a := str[0]
	b := str[1]
	fmt.Printf("a=%c\n", a)
	fmt.Printf("b=%c\n", b)

	var byteSlice []byte
	byteSlice = []byte(str)
	byteSlice[0] = 'M'
	str = string(byteSlice)
	fmt.Printf("str=%s\n", str)
}

func StringTest2() {

	str2 := "Hello,世界"

	var str2ByteSlice []byte
	str2ByteSlice = []byte(str2)

	var str2RuneSlice []rune
	str2RuneSlice = []rune(str2)

	fmt.Printf("strLen=%v\n",len(str2) )
	fmt.Printf("str2ByteSlice=%v\n",len(str2ByteSlice) )
	fmt.Printf("str2RuneSlice=%v\n",len(str2RuneSlice) )

	for i:=0;i<len(str2RuneSlice);i++{
		var b = str2RuneSlice[i]
		fmt.Printf("str2RuneSlice[%d]=%c\n" ,i,b)
	}



}