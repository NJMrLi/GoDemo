package GoroutineDemo


import (

"fmt"

"runtime"

)



func say(s string) {

	for i := 0; i < 2; i++ {

		runtime.Gosched()

		fmt.Println(s)

	}

}



func GoschedDemo() {

	go say("world")

	say("hello")

}