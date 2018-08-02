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
//import "fmt"
//
//func main() {
//	a := 77
//	fmt.Println(a)
//	fmt.Println(&a)
//
//	var b *int = &a
//	fmt.Println(b)
//	fmt.Println(*b)
//	*b=75
//	fmt.Println(a) //we changed the value here by call by reference
//}
//
//Loops
//for loop to print a series of number
//import "fmt"
//
//func main(){
//	for i:=3;i<=10;i++ {
//		fmt.Println(i)
//	}
////	for j:=3;j<=10;++j{
////		fmt.Println(j)
////	}
//}
//
//Nested For loops
//import "fmt"
//
//func main(){
//	for i:=0 ;i<5; i++{
//		fmt.Print("\t")
//		for j:=0;j<=5;j++{
//			fmt.Print(i ,"-" ,j,"\n")
//		}
//	}
//}
//
//While like loop with for loop
//import "fmt"
//
//func main(){
//	i:= 10 //initialize
//	for i<100 {
//		fmt.Println(i)
//		i++
//	}
//}
//
//Infinite one
//import "fmt"
//
//func main(){
//	i:=0
//	for i<10 {
//		i++
//		fmt.Println(i)
//	}
//}
//
//Lets use break with conditions
//import "fmt"
//
//func main(){
//	i:=0
//	for {
//		fmt.Println(i)
//		if i>=11{ //here 11 will be printed before the condition is checked
//			break;
//		}
//		i++
//	}
//}
//
//Program to get user 's input and show Odd/Even
//import "fmt"
//
//func main() {
//	var check int
//	fmt.Println("Enter the number")
//	fmt.Scan(&check)
////	i := 0
//	if check%2 == 0 {
//		fmt.Println(check, " is even")
//	}	else{
//		fmt.Println(check, " is odd")
//	}
//}
//
//Program to print only odd no from series of number
//import "fmt"
//
//func main(){
//	var num1,num2 int
//	fmt.Println("Enter The range below to get all the odd numbers in it")
//	fmt.Println("Enter First Number")
//	fmt.Scan(&num1)
//	fmt.Println("Enter Second Number")
//	fmt.Scan(&num2)
//	i:=num1 -1 //-1 so that that first num is not left out from the loop
//	for i<num2 {
//		i++
//		if i%2==0{
//			continue
//		}
//		fmt.Println("The odd num is",i)
//	}
//}
//
//Program to print unicode of string/Number //consider going through `` and rune conversion once again if not getting from here
//import "fmt"
//
//func main(){
//	var char11 string
//	fmt.Println("Enter the character you want unicode of")
//	fmt.Scan(&char11)
////	fmt.Println("The unicode for your entered value is as follow",[]byte(char11))
//	fmt.Printf("%v \n %v","The unicode for your entered value is",[]byte(char11))
//}
//
//Switch Statements and its "fallthrough"
//import "fmt"
//
//func main(){
//	var i string
//	fmt.Println("Write from the names below")
//	fmt.Println("ravi,ram,raju,sam")
//	fmt.Scan(&i)
//	switch{
//		case len(i)==2:
//			fmt.Println("This is selected if the string length is of two words")
//		case i=="ravi",i=="sam":
//			fmt.Println("hello ravi or sam")
//		case i=="ram":
//			fmt.Println("hello ram")
//		case i=="raju":
//			fmt.Println("hello raju")
//		default:
//			fmt.Println("Well the name is not in the directory")
//	}
//}
//
//Exercises question 6
//import "fmt"
//
//func main() {
//	for i := 1; i <= 100; i++ {
//		if i%3 == 0  && i%5 == 0 {//or we can do i%15
//			fmt.Println(i,"FizzBuzz")
//		} else if i%3 == 0 {
//			fmt.Println(i,"Fizz")
//		} else if i%5 == 0 {
//			fmt.Println(i,"Buzz")
//		}else {
//			fmt.Println(i)
//		}
//	}
//}
//
//Exercise question 7
//import "fmt"
//
//func main() {
//	sum := 0
//	for i := 1; i < 1000; i++ {
//		if i%3 == 0 || i%5 == 0 {
//			sum = sum + i
//			//			fmt.Println(sum)
//		}
//	}
//	fmt.Println(sum)
//}
//
//Returning a function
//import "fmt"
//
//func main() {
//	fmt.Println(greet("abhijit", "thapa"))
//}
//func greet(fname, lname string) string { //returns single parameter
//	return fmt.Sprint(fname, "\t", lname)
//}
//
//Returning a function with two parameters
//import "fmt"
//
//func main() {
//	fmt.Println(greet("Abhijit", "Thapa"))
//}
//func greet(fname, lname string) (string, string) { //Two Parameters
//	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
//}
//
//Get The average of numbers //Recheck Video 69////////////////////////////////////////
import "fmt"

func main() {
	data:= []int{25, 42, 73, 43, 64, 84}
	n := average(data...)
	fmt.Println(n)
}
func average(av ...int) int {
	fmt.Println(av) //shown in slice
	fmt.Printf("%T \n", av) //shown in slice
	var total int
	for _, tot := range av { //range loops over list of items
		total += tot //shorthand for  loop
	}
	return total / len(av) //returns total
}
