package InterfaceDemo

import (
	"fmt"
	"reflect"
)

type eat interface {
	eat() (string)
}

type Bird struct {
	Name string
}

func (bird Bird) eat() string {
	return "Bird:" + bird.Name + " eat"
}

func print(e eat) {
	fmt.Println(e.eat())
}

func ITableDemo() {


	bird1 := Bird{Name:"Big"}
	bird2 := new(Bird)
	bird2.Name = "Small"
	fmt.Println(reflect.TypeOf(bird1).String())
	fmt.Println(reflect.TypeOf(bird2).String())

	print(bird1)
	print(bird2)

	//下面就是使用接口来接收具体的值
	//参考《Go语言实战》 5.4.2 实现中的两幅图理解一下
	var eat1 eat
	eat1 = bird1  //传递的是值 实体值赋值后接口值
	print(eat1)

	var eat2 eat
	eat2 = bird2 //传递的是地址 实体指针赋值后接口值
	print(eat2)


}
