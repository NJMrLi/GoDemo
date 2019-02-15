package FuncDemo

import "fmt"

type Student struct {
	Name string
	Age  int
}

func GetStudentName(student Student) (name string) {
	//返回值
	return student.Name
}

//Student类的方法 使用值接收者实现了一个方法
func (student Student) GetStudentName() (string) {
	return student.Name
}

//Student类的方法 使用指针接收者实现了一个方法
func (student *Student) ChangeStudentName(name string) {
	student.Name = name
	fmt.Println("方法执行之后name", student.Name)

}

//Student类的方法 使用指针接收者实现了一个方法
func (student Student) ChangeStudentNameByValue(name string) {
	student.Name = name
	fmt.Println("方法执行之后name", student.Name)
}

func DiffFuncAndMethod() {

	bigOrange := Student{
		Name: "BigOrange", Age: 18,
	}

	bigApple := Student{
		Name: "BigApple", Age: 20,
	}

	//使用函数获取学生名称
	name1 := GetStudentName(bigOrange)
	name2 := GetStudentName(bigApple)

	fmt.Println("========通过传地址ChangeName之后==========")
	fmt.Println("方法执行之前name", name1)
	bigOrange.ChangeStudentName("BigBanana")
	name1 = bigOrange.GetStudentName()
	fmt.Println("方法返回之后Name", name1)

	fmt.Println("========通过传值ChangeName之后===========")
	fmt.Println("方法执行之前name", name2)
	bigApple.ChangeStudentNameByValue("BigPear")
	name2 = bigApple.GetStudentName()
	fmt.Println("方法返回之后Name", name2)

}
