package SliceDemo

import "fmt"

func GetSliceCapAndLength(){
	slice1:=[]int{10,20,30,40,50}
	slice2:=[]int{10,20,30,40,50,60,70,80}

	var len1 = len(slice1)
	var len2 = len(slice2)
	var cap1 = len(slice1)
	var cap2 = len(slice2)

	fmt.Println("len1",len1)
	fmt.Println("cap1",cap1)
	fmt.Println("len2",len2)
	fmt.Println("cap2",cap2)

	var newslice1 = slice1[1:3]
	var newslice2 = slice1[1:4]
	var newslice3 = slice2[1:4]

	var newlen1 = len(newslice1)
	var newlen2 = len(newslice2)
	var newlen3 = len(newslice3)
	var newcap1 = cap(newslice1)
	var newcap2 = cap(newslice2)
	var newcap3 = cap(newslice3)

	fmt.Println("newlen1",newlen1)
	fmt.Println("newcap1",newcap1)
	fmt.Println("newlen2",newlen2)
	fmt.Println("newcap2",newcap2)
	fmt.Println("newlen3",newlen3)
	fmt.Println("newcap3",newcap3)
}

func GetAppendSlice(){

	array:=[6]int{0,1,2,3,4,5}
	myslice:=array[0:2:4]
	myslice2:=array[2:4:5]

	var lens1 = len(myslice)
	var caps1 = cap(myslice)
	var lens2 = len(myslice2)
	var caps2 = cap(myslice2)


	fmt.Println("长度:",lens1)
	fmt.Println("容量:",caps1)
	fmt.Println("长度:",lens2)
	fmt.Println("容量:",caps2)

	fmt.Println("修改前myslice[0]",myslice[0])
	fmt.Println("修改前myslice[1]",myslice[1])
	fmt.Println("修改前myslice2[0]",myslice2[0])
	fmt.Println("修改前myslice2[1]",myslice2[1])

	myslice = append(myslice,111,222)

	fmt.Println("修改后myslice[0]",myslice[0])
	fmt.Println("修改后myslice[1]",myslice[1])
	fmt.Println("修改后append数据 myslice[2]",myslice[2])
	fmt.Println("修改后append数据 myslice[3]",myslice[3])
	fmt.Println("修改后myslice2[0]",myslice2[0])
	fmt.Println("修改后myslice2[1]",myslice2[1])
}