//hello world
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("42")
//}
//print numbers
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println(42)
//}
//Print values in different numerical forms
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Printf("%d - %b", 42, 42) prints numeber in decimal and binary form
//	fmt.Printf("%x -%X - %o \n", 42,42,42)  %X   for bigger hex letter
//	fmt.Printf("%d - %b \n", 42, 42) \n   for next line
//	fmt.Println("%d \t - %b \t", 42, 42)  nothing happens \t gives tab space
//	fmt.Print("%d - %b", 42, 42)nothing   happens shown same as here
//}
// loops
package main

//import (
//	"fmt"
//	"github.com/abhi-th/gothird/gopkg1"
//	"github.com/abhi-th/gothird/stringutil"
//)

//	var a int= 5 //package scope also refer D:\Go\GoCode\src\github.com\abhi-th\GolangTraining-master\03_variables\03_less-emphasis\07_all-together
//func main() {
//get the sum
//	sum := 0
//	for i := 0; i < 20; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//simple loops
//	for i:=0; i< 200; i++{
//		fmt.Printf("%d \t %b \t %#x \t %#o \n",i,i,i,i)
//	}
//	for i:=50;i<140;i++{
//		fmt.Printf("%d \t %b \t %x \t %o %q \n",i,i,i,i,i)
//	}
//Scope
//fmt.Println(gopkg1.MyName1,gopkg1.MyName2,gopkg1.MyName3)
//reverse
//	fmt.Println(stringutil.Reverse("!oG ,olleH"))
//	fmt.Println(stringutil.MyName)
//	fmt.Println(winniepooh.BearName)
//
// %v shows value(given & default) %T shows type
//	var a int = 10
//	var b string
//	var c float64
//	var d byte
//	var e bool
//
//	fmt.Printf("%v \n",a)
//	fmt.Printf("%T \n",b)
//	fmt.Printf("%T \n",c)
//	fmt.Printf("%T \n",d)
//	fmt.Printf("%T \n",e)
//
//Declaration and initialization
//	var a int= 7
//	b:= "Iron Network"
//	c:="0.6"
//	d:= true
//	fmt.Printf("%T \t",a)
//	fmt.Printf("%T \t",b)
//	fmt.Printf("%T \t",c)
//	fmt.Printf("%T \t",d)
//	fmt.Printf("%v \t",a)
//	fmt.Printf("%v \t",b)
//	fmt.Printf("%v \t",c)
//	fmt.Printf("%v \t",d)
//
//Exercises from resource folder
//1
//	var message string="Hello World through variable"
//	var a,b,c int = 1,2,43
//	d,e,f:=4 ,false, 01010
//	var g,h,i = 1,2,43
//
//	fmt.Println(message,a,b,c,d,e,f,g,h,i)
//	fmt.Printf("%T",f) // why this is showing int value here along with f = 65
//	d = "stored in d" // declaration above; assignment here; package scope(refer in the less_emphasis alltogethre folder)
//	var e = 42        // function scope - subsequent variables have func scope:
//	f := 43
//	g := "stored in g"
//	h, i := "stored in h", "stored in i"
//	j, k, l, m := 44.7, true, false, 'm' // single quotes
//	n := "n"                             // double quotes
//	o := `o`                             // back ticks
//	fmt.Println("e - ", e)
//	fmt.Println("f - ", f)
//	fmt.Println("g - ", g)
//	fmt.Println("h - ", h)
//	fmt.Println("i - ", i)
//	fmt.Println("j - ", j)
//	fmt.Println("k - ", k)
//	fmt.Println("l - ", l)
//	fmt.Println("m - ", m)
//	fmt.Println("n - ", n)
//	fmt.Println("o - ", o)
//
//}

//30 july Scope 1
//import "fmt" //scope of import is file level as it can only be used in this file only
//
//var x int = 42 //scope here is package level, as anything relating to this pakage main can use this variable but then this should be in capitalized form i.e X and not x
//
//func main() {
//	fmt.Println(x)
//	foo()
//	y:=10
//	fmt.Println(y)// remember the ordering of variables here really matters
//}
//func foo() {
//	fmt.Println(x)
////	fmt.Println(y) // scope of y is block level
//}

////Closure
//
//import "fmt"
//
//var x = 0
//
//func increment() int {
//	x++
//	return x
//}
//func main() {
//	fmt.Println(x)
//	fmt.Println(increment())
//	fmt.Println(x)
//	fmt.Println(increment())
//	fmt.Println(x)
//}
//2. CLosure
//import "fmt"
//
//func main(){
//	x:=42
//	{
//	fmt.Println(x)
//	y:= "The credit belongs to those in the ring"
//	fmt.Println(y)
//	}
////	fmt.Println(y)
//	fmt.Println(x)
//}
//3. Closure //func expression
//import "fmt"
//
//func main(){
//	x:=0
//	increment:=func() int{ //func expression //anonymous function(a function without name)
//		x++
//		return x
//	}
//	fmt.Println(increment())
//	fmt.Println(increment())
//}
//4.CLosure Returning a Function
//import "fmt"
//
//func wrapper() func() int {
//	x := 0
//	return func() int {
//		x++
//		return x
//	}
//}
//func main() {
//	increment := wrapper()
//	fmt.Println(increment())
//	fmt.Println(increment())
//}
//5.Blank Identifier
//import (
//	"net/http"
//	"io/ioutil"
//	"fmt"
//)
//func main(){
//	res, _ :=http.Get("http://bizhub.theironnetwork.org")
//	page,_ :=ioutil.ReadAll(res.Body)
//	res.Body.Close()
//	fmt.Printf("%s",page)
//}
//
//Constants 1
//import "fmt"
//
//const q = 42
//
//func main() {
//	const p = "lives & taxes"
//	fmt.Println("p-", p)
//	fmt.Println(p, q)
//	fmt.Println("q-", q)
//}

//Constants 2 Multiple constants
//import "fmt"
//
//const (
//	Pi       = 3.14
//	Language = "Go"
//)
//
//func main() {
//	fmt.Print(Pi)
//	fmt.Println(Language)
//}
//
//Constants 3 Iota
//import "fmt"
//
//const (
//	A = iota
//	B = iota
//	C = iota
//)
//const (
//	D = iota
//	E
//	F
//)
//const (
//	G = iota
//	H
//	I
//)
//func main() {
//	fmt.Println(A, B, C, D, E, F, G, H, I)
//}
////this increments value .Also we need to declare atleast once "iota" with the top most constant to make the value repeat.Also the value gets redefined if we redeclare constants
//
//Constants 4 iota
//import "fmt"
//
//const (
//	_  = iota
//	KB = 1 << (iota * 10)//<< here means bit shifting (1X10)means bit shifting is 2^10 times(in binary )
//	MB = 1 << (iota * 10)//<< here means bit shifting (2X10)means bit shifting is 2^20 times
//	GB = 1 << (iota * 10)
//	TB = 1 << (iota * 10)
//)
//
//func main() {
//	fmt.Printf(" binary value - %b\t\t\t\t\t", KB)
//	fmt.Printf(" decimal value - %d\n", KB)
//	fmt.Printf(" binary value - %b\t\t\t\t", MB)
//	fmt.Printf(" decimal value - %d\n", MB)
//	fmt.Printf(" binary value - %b\t\t\t", GB)
//	fmt.Printf(" decimal value - %d\n", GB)
//	fmt.Printf(" binary value - %b\t", TB)
//	fmt.Printf(" decimal value - %d\n", TB)
//}
//
//Memory Addresses 1
//import "fmt"
//
//func main() {
//	x := 88
//	fmt.Println(&x)//hex value of address(small post box)
//	fmt.Printf("%d \n",&x)//decimal value of address(small post box)
//}
//Memory address 2 //Enter input
//import "fmt"
//
//const metersToYards= 1.09361
//
//func main() {
//	var meters float64
//	fmt.Print("Enter the meters you swam - ")
//	fmt.Scan(&meters)
//	yards := meters * metersToYards
//	fmt.Println("You swam for about ", yards, " yards")
//}
import "fmt"

func main() {
	a := 77
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
	*b=75
	fmt.Println(a) //we changed the value here by call by reference
}
