package main

import "fmt"
//////////////////////////////分支测试
import (
	"fmt"
	//"runtime"
)


var dstr string= "sdf"
var bstr bool=true
var cstr="ttt"
//结构体
type book struct {
	title string
	id int
	makstr string
}
func init() {
	//首先 init 在 main 函数
	fmt.Println(cstr+ "  init 输出 hello\nworld 变量："+dstr);fmt.Println("init 输出 helloworld2 变量："+dstr)
	fmt.Println(bstr)
}

func main() {
	fmt.Println("helloworld")
	/*//测试1
	fmt.Println("main hello world 变量："+dstr)
	fmt.Println("dstr变量内存地址：",&dstr)
	var  dint=123
	var dint2 *int
	dint2=&dint
	fmt.Println("123内存地址",&dint)
	fmt.Println("dint2值",dint2)
	fmt.Println("*dint2值",*dint2)

	var book1 book
	book1.id=1
	book1.makstr="备注"
	book1.title="go语言"
	fmt.Println("book1",book1)
	fmt.Println("book1",book1.id)
	var book2 book
	book2.id=2
	book2.makstr="备注2"
	book2.title="c++语言"
	printBook1(book2)

//	go say("world")//与 下面的go say world1 一起执行
	//say("hello")

//	go say2("world1")//为什么 go say2 下面还要有个 say2方法 只有一个go say2不执行
	go say2("hello")
 	say2("hello1")
*/

}/*
func printBook1(booktemp book ){
	fmt.Println("传入book",booktemp)
}
func say(s string){
	for i:=0;i<5;i++{
		runtime.Gosched() //如果注释掉 只能顺序执行
		fmt.Println(s,":",i)
	}
}
func say2(s string){
	for i:=0;i<5;i++{
		runtime.Gosched()
		//runtime.GOMAXPROCS(2)
		fmt.Println(s,":",i)
	}
}
*/