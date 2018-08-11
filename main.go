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
//import "fmt"
//
//func main() {
//	data := []float64{25, 42, 73, 43, 64, 84}
//	n := average(data...)
//	fmt.Println(n)
//}
//func average(av ...float64) float64 {
//	fmt.Println(av)         //shown in slice
//	fmt.Printf("%T \n", av) //shown in slice
//	var total flaot64
//	for _, tot := range av { //range loops over list of items
//		total += tot //shorthand for  loop
//	}
//	return total / float64(len(av)) //returns total
//}
//
//Edited one(same slice typed) //This works same and depends on the situation for use
//import "fmt"
//
//func main() {
//	data := []float64{25, 42, 73, 43, 64, 84}
//	n := average(data)
//	fmt.Println(n)
//}
//func average(av []float64) float64 {
//	fmt.Println(av)         //shown in slice
//	fmt.Printf("%T \n", av) //shown in slice
//	var total float64
//	for _, tot := range av { //range loops over list of items
//		total += tot //shorthand for  loop
//	}
//	return total / float64(len(av)) //returns total
//}
//
//Reviews Closure and func expression
//function without closure and func expression
//import "fmt"
//
//func greeting(){
//	fmt.Println("Hello World")
//}
//
//func main(){
//	greeting()
//}
//function with funcexpression //only way for func inside function
//import "fmt"
//
//func main(){
//	var x int
//	increment:=func() int{
//		x++
//		return x
//	}
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//}
//another way for func expression
//import "fmt"
//var x int
//func makeincrementor() func() int{
//	return func() int{
//		x++
//		return x
//	}
//}
//func main(){
//	increment:= makeincrementor()
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//	fmt.Println(increment())
//}
//
//CLosure same previous examples
//
//Callback
//import "fmt"
//
//func values(numbers []int, callback func() int) {
//	for _, n := range numbers {
//		callback(n)
//	}
//}
//func main() {
//	values([]int{2, 4, 6, 3, 5}, func(n int) {
//		fmt.Println(n)
//	})
//}
//
//Call back example 2
//import "fmt"
//
//func main() {
//	ab := filter([]int{2, 43, 54, 5, 3}, func(n int) bool {
//		return n > 1
//	})
//	fmt.Println(ab)
//}
//
//func filter(numbers []int,callback func(int) bool) []int{
//	var ab []int
//	for _,n:=range numbers{
//		if callback(n){
//			ab=append(ab,n)
//		}
//	}
//	return ab
//}
//
//Recursion
//import "fmt"
// var x int
//func factorial(x int)int{
//	if x==0{
//		return 1
//	}
//	return x * factorial(x-1)
//}
//
//func main(){
//	fmt.Println("Enter the number")
//	fmt.Scan(&x)
//	fmt.Println("You factorial value for this no is ",factorial(x))
//}
//
//import "fmt"
//
//func main(){
//	fmt.Println(factorial(4))
//}
//
//func factorial(x int)int{
//	if x==0{
//		 return 1
//	}
//	return x * factorial(x-1)
//}
//
//DEfer // sexecutes the func rigth before exit
//import "fmt"
//
//func main() {
//
// 	defer world()
//	hello()
//}
//func hello() {
//	fmt.Println("hello ")
//}
//func world() {
//	fmt.Println("World")
//}
//Another example of defer file with open and close//Limit reader
//import (
//	"os"
//)
//
//func main() {
//	src, err := os.Open("src.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer src.Close()
//
//	dst, err := os.Create("dst.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer dst.Close()
//
//	bs := make([]byte, 5)
//	src.Read(bs)
//	dst.Write(bs)
//}
//
// this is a limit reader
// we limit what is read
// see io.ReadFull for func similiar to (*File)Read
//
//Exercises
//Exercise 1
//import "fmt"
//
//func checkeven(n int) /*this is parameter*/ (int, bool) /*and these are returns */ {
//	return n / 2, n%2 == 0 /*two args for return*/
//}
//
//func main() {
//	a, b := checkeven(5)
//	fmt.Println(a, b)
//}
//
//Exercise 2
//import "fmt"
//
//func main() {
//	checkeven := func(n int) (int, bool) {
//		return n / 2, n%2 == 0
//	}
//	a, b := checkeven(5)
//	fmt.Println(a, b)
//}
//
//Exercise 3
//import "fmt"
//
//func main(){
//	greatest:=max(3,7,5,9,8,78,3)
//	fmt.Println(greatest)
//}
//func max(numbers ...int)int{
//	var largest int
//	for _,v := range numbers{
//		if v > largest{
//			largest=v
//		}
//	}
//	return largest
//}
//Exercise 4
//import "fmt"
//
//func main() {
//	if (true && false) || (false && true) || !(false && false) {
//		fmt.Println("true")
//	} else {
//		fmt.Println("false")
//	}
// //or below solution
//fmt.Println((true&&false)||(false&&true)||!(false&&false))
//}
//Exercise 5
//import "fmt"
//
//func main() {
//	foo(1, 2)
//	foo(1, 2, 3)
//
//	aSlice := []int{1, 2, 3, 4, 5}
//	foo(aSlice...)
//
//	foo()
//}
//func foo(numbers ...int) {
//	fmt.Println(numbers)
//}
//
//Copy Slice
//import "fmt"
//
//func main(){
//	slice1:=[]int{1,2,3,4,5}
//	slice2:= make([]int,2)
//	copy(slice2,slice1)
//	fmt.Println(slice2)
//}
//
//Question 2
//fibonaci series
//import "fmt"
//
//func fibonacci(n int) int {
//	if n <= 1 {
//		return 1
//	}
//	return fibonacci(n-1) + fibonacci(n-2)
//}
//func main() {
//	n := 0
//	var sum int
//	for i := n; fibonacci(i) < 4000000; i++ {
//		if fibonacci(i)%2 == 0 {
//			//			fmt.Println("Even value check ", fibonacci(i))
//			sum = sum + fibonacci(i)
//		}
//		//		fmt.Println("overall values till loop ",i, fibonacci(i))
//	}
//	fmt.Println("The sum of even valued terms is", sum)
//}
//import (
//	"fmt"
//		//"math"
//)
//
//func main() {
//	var pri float64
//	var i,n float64
//	fmt.Println("Enter no to Print factorS of it")
//	fmt.Scan(&pri)
//	//	pri =math.Sqrt(pri)
//	for i = 2; i < pri; i++ {
//		if pri%i == 0 {
//				for n := 2; n < 10 ; n++ {
//					if i%n != 0 {
//						i= math.Sqrt(i)
//					}
//				}
//					fmt.Print(i, "\t")
//		}
//	}
//}
//import "fmt"
//import "math"
//
//func main(){
//	var num int
//	fmt.Println("Enter your number")
//	fmt.Scan(&num)
//    for t := 2; t < num; t++ {
//        if num%t == 0 {
////			  fmt.Println(i,"ex")
////			  for n:=2 ;n< 10;n++{
////				  if i%n==0{
//////					  fmt.Println("see",n,i)
////					 continue
////				  }
////				  fmt.Println("see1",i)
////			  }
//			  fmt.Println(IsPrime(t))
//        }
//    }
//}
//func IsPrime(value int) int {
//    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
//        if value%i == 0 {
//			  continue
//        }
//    }
//    return value
//}
//import "fmt"
//
//func main() {
//	var num,largest int
//	myslice:= make([]int,0,100)
//	fmt.Print("Enter Number Here : ")
//	fmt.Scan(&num)
//	fmt.Print("Prime Factors are : ")
//	for i := 2; i <= num; i++ {
//		if num%i == 0 {
//			//fmt.Println(i)
//			s := notPrime(i)
//			if !s {
//				fmt.Print(i,"\t")
//				myslice=append(myslice,i)
//			} else {
//				continue
//				//    fmt.Println("Not Prime Factor =", i)
//			}
//		}
//	}
////	fmt.Println(myslice)
//	for _,v:=range myslice{
//		if v>largest{
//			largest=v
//		}
//	}
//	fmt.Println("\n","The largest Prime Factor is : ",largest)
//}
//
//func notPrime(num int) bool {
//	var notprime bool
//	for i := 2; i <= num/2; i++ {
//		if num%i == 0 && num > 2 {
//			notprime = true
//			break
//		} else {
//			notprime = false
//		}
//	}
//	return notprime
//}

//
//Array 1
//
//import ("fmt")
//
//func main(){
//	var x[58] string
//	fmt.Println(x)
//	fmt.Println(len(x))
//	fmt.Println(x[0])
//	x[42]=777
//	for i:=65;i<=122;i++{
//		x[i - 65]= string(i)
//	}
//	fmt.Println(x)
//	fmt.Println(len(x))
//	fmt.Println(x[0])
//}
//
//Array 2
//
//import "fmt"
//
//func main(){
//	var x[58] string
//	for i:=65; i<=122;i++{
//		x[i-65] =string(i)
//	}
//	fmt.Println(x)
//	fmt.Println(len(x))
//	fmt.Println(x[22])
//}
//
//Array 3
//
//import "fmt"
//
//func main(){
//	var x[256] int
//	for i:=0;i<256;i++{
//		x[i]= i
//	}
//	for item,value:=range x{
//		fmt.Println(item,value)
////		fmt.Println(item,value)
//	}
//}
//Slices
//
//import "fmt"
//
//func main(){
//	mySlice :=make([]int ,0,5)
//	fmt.Println("------------------------")
//	fmt.Println(mySlice)
//	fmt.Println(len(mySlice))
//	fmt.Println(cap(mySlice))
//	fmt.Println("------------------------")
//	for i:=0;i<80;i++{
//		mySlice= append(mySlice,i)
////		fmt.Println(" len ",len(mySlice)," Capacity ",cap(mySlice),mySlice)
//	}
//		fmt.Println(mySlice)
//}
//
//Slice Example 2
//import "fmt"
//
//func main(){
//	employee:=[]string{
//		"name1",
//		"name2",
//		"name3",
//		"name4",
//		"name5",
//	}
//	//way 1
//	for i,v:=range employee{
//		fmt.Println(i,v)
//	}
//	//way 2
//	for i:=0;i<len(employee);i++{
//		fmt.Println(employee[i])
//	}
//}
//Slice Example 3
//import "fmt"
//
//func main(){
//	records:= make([][]string,4,10)
//	//student
//	student1:=make([]string,4)
//	student1[0]="raman"
//	student1[1]="verma"
//	student1[2]="22"
//	student1[3]="44"
//	records= append(records,student1)
//	//student
//	student2:=make([]string,4)
//	student2[0]="ajay"
//	student2[1]="sharma"
//	student2[2]="24"
//	student2[3]="88"
//
//	records= append(records,student1)
//	fmt.Println(records)
//}
//
//slice example 4
//import "fmt"
//
//func main(){
//	transactions:=make([][]int,0,3)
//	for i:=0; i<3;i++{
//		transaction:=make([]int,0)
//		for j:=0;j<4;j++{
//			transaction=append(transaction,j)
//		}
//		transactions=append(transactions,transaction)
//	}
//	fmt.Println(transactions)
//
//}
//Question 5
//import "fmt"
//
//var i, j int
//
//func main() {
//Loop:
//	for i = 1; i > 0; i++ {
//		for j = 1; j <= 20; j++ {
//			if i%j == 0 {
//				if j == 20 {
//					fmt.Println(i)
//					break Loop
//				} else {
//					continue
//				}
//			}else{
//				break
//			}
//		}
//	}
//}
//import "fmt"
//
//var k,l int
//func main(){
//	j:=1
//	ps:=1
//	for i:=1;i<=100;i++{
//		j=i*i
////		fmt.Println("product",j)
//		k=k+j
//		l=l+i
//	}
//	ps=l*l
//	fmt.Println("product of sum",ps)
//	fmt.Println("sum of products",k)
//	fmt.Println("the difference between the above values are",ps-k)
//
//}
//
//Maps Example Part 1
//Example 1
//
//import "fmt"
//
//func main(){
//	m:=make(map[string]int) //map [key type] value type
//
//	m["k1"]=7
//	m["k2"]=13
//	m["k3"]=14
//
//	fmt.Println(m)
//}
//
//Example 2
//import "fmt"
//
//func main(){
////	var greet= map[int]string{ \\1st way
////		1:"hello",
////		2:"hi",
////	}
//	var greet= make(map[int]string)
//	greet[1]="hello"
//	greet[2]="hello"
//	greet[3]="Bonjour"
//	greet[4]="Dias"
//	delete(greet,4) //delete(m,k1)
//	fmt.Println(greet)
//}
//https://www.golang-book.com/books/intro/6
//Conditionalize Map
//import "fmt"
//
//func main(){
//	greet := map[int]string{
//		1:"hi",
//		2:"bonjour",
//		3:"Hello",
//		4:"Morning",
//	}
//	fmt.Println(greet)
//	delete(greet,2)
//
//	if _, exists:=greet[2];exists{     //Condition to chexk if the value exists or not//COMMA OK
//		fmt.Println("exist")
//		fmt.Println(greet)
//	}else {
//		fmt.Println(greet)
//		fmt.Println("doesn't exist")
//	}
//}
//
//Range With Maps
//
//import "fmt"
//
//func main(){
//	greet:=map[int]string{
//		1:"hi",
//		2:"bonjour",
//		3:"Hello",
//		4:"Morning",
//	}
//
//	for k,v:=range greet{
//		fmt.Println(k,v)
//	}
//}
//Example
//
// hash table Step 1
//import "fmt"
//
//func main(){
////	letter:='A' //Singe quotes for runes
////	letter:=rune("Abhijit"[0])//for "" string give me the first position and turn it into rune
////	fmt.Println(letter)
////	fmt.Printf("%T \n")
////	for i:=65;i<=122;i++{
////		fmt.Println(i," - ",string(i)," - ",i % 12)
////	}
//	n:=HashBucket("Go",12)
//	fmt.Println(n)
//}
//
//func HashBucket(word string,buckets int) int{
////	letter:=rune(word[0]) //but for this we need to make the return type as int 32 els eit would give type mismatch error
//	letter:= int(word[0])
//	bucket:= letter % buckets
//	return bucket
//}
////Read/scan Function
//import (
//	"fmt"
//	"bufio"
//	"os"
//	"strings"
//)
//
//func main(){
//	const input="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Alias tenetur, aspernatur sit eligendi in cumque tempore odit est asperiores, beatae a omnis. Esse ab non blanditiis odit unde quas, nemo iure necessitatibus dolores ipsa culpa sequi quia sapiente. In, assumenda consequuntur. Sequi rerum velit consequuntur sit suscipit ratione. Id voluptatum corrupti cum veritatis natus rerum voluptatem, sequi quibusdam. Dolorum consequatur eveniet maiores nostrum saepe officiis repellendus autem, magni ut reprehenderit aliquid vel laborum quibusdam, ducimus quisquam nulla qui iusto consequuntur deserunt possimus. Quas ratione nostrum illo, eos culpa. Ad quas qui quia iste eveniet repellendus voluptatum, tenetur molestiae ipsam perspiciatis."
//	//newscanner provides a pointer to the scanner
//	scanner := bufio.NewScanner(strings.NewReader(input))
//	//set the split function for the scanning operation.
//	scanner.Split(bufio.ScanWords)
//	//count the words.
//	for scanner.Scan(){
//		fmt.Println(scanner.Text())
//	}
//	if err:=scanner.Err();err!=nil{
//		fmt.Fprintln(os.Stderr,"reading input:",err)
//	}
//}
//Get the book and read all
//
//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//)
//
//func main(){
//	res,err:= http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
//	if err != nil{
//		log.Fatal(err)
//	}
//	bs,err := ioutil.ReadAll(res.Body)//put in byteslice
//	res.Body.Close()
//	if err != nil{
//		log.Fatal(err)
//	}
//	fmt.Printf("%s",bs)//printed as a string
//}
//
//Struct
//
//import "fmt"
//
//type person struct{
//	first string
//	last string
//	age float64
//}
//
//func main(){
//	p1:=person{"James","Bond",20.9}
//	p2:=person{"Sukhpal","Singh",14}
//	fmt.Println(p1.first,p1.last,p1.age)
//	fmt.Println(p2.first,p2.last,p2.age)
//}
//
//User Defined Types
//import "fmt"
//
//type  foo int
//
//func main(){
//	var myage foo
//	myage=44
//	fmt.Printf("%T %v \n",myage,myage) //type here is main.foo
//
//	var yourage int
//	yourage=29
//	fmt.Printf("%T %v \n",yourage,yourage)
//
//	sum:=int(myage)+yourage
//	fmt.Println(sum)
//}
// We used shorthand notation:
// to create a variable named p1 of type person
// to create a variable named p2 of type person
// We initialized those variables with specific values
// We used the short variable declaration operator with a struct literal to initialize
// ----------------------------------------
// here is how we talk about structs:
// -- user defined type
// -- we declare the type
// -- the type has fields
// -- the type can also have "tags"
// ----  we haven't seen this yet
// -- the type has an underlying type
// ---- in this case, the underlying type is struct
// -- we declare variables of the type
// -- we initialize those variables
// ---- initialize with a specific value, or
// ---- or, initiliaze to the zero value
// -- a struct is a composite type
// ----------------------------------------
// Bill Kennedy:
// Go allows us the ability to declare our own types.
// Struct types are declared by composing a fixed set of unique fields together.
// Each field in a struct is declared with a known type.
// This could be a built-in type or another user defined type.
// Once we have a type declared, we can create values from the type
// When we declare variables, the value that the variable represents is always initialized.
// The value can be initialized with a specific value or it can be initialized to its zero value
// For numeric types, the zero value would be 0; for strings it would be empty;
// and for booleans it would be false.
// In the case of a struct, the zero value would apply to all the different fields in the struct.
// Anytime a variable is created and initialized to its zero value, it is idiomatic to use the keyword var.
// Reserve the use of the keyword var as a way to indicate that a variable is being set to its zero value.
// If the variable will be initialized to something other than its zero value,
// then use the short variable declaration operator with a struct literal
//
//Pattern making
//import "fmt"
//
//func main(){
//	for i:=5 ;i>=1;i--{
//		for j:=1;j<=i;j++{
//			fmt.Print(" *")
//		}
//		fmt.Printf("\n")
//	}
//}
//
//type overridding
//import "fmt"
//
//type person struct{
//	first string
//	last string
//	age int
//}
//
//type doubleZero struct{
//	person
//	first string
//	spy bool
//}
//func main(){
//	p1:=doubleZero{
//		person:person{
//			first: "James" ,
//			last:"bond",
//			age:22,
//		},
////		first:"john",
//		spy:true,
//	}
//	p2:=doubleZero{
//		person:person{
//			first:"Asif",
//			last:"Ali",
//			age:28,
//		},
//		first: "Example",
//		spy:false,
//	}
//	fmt.Println(p1.first,p1.last,p1.age) //here overriding is done
//	fmt.Println(p1.person.first,p1.person.last)
//	fmt.Println(p2.first,p2.last,p2.age)
//}
//
//Method Override
//import "fmt"
//
//type student struct{
//	first string
//	last string
//	rollno int
//}
//type classrep struct{
//	student
//	classrep bool
//	marks int
//}
//
//func (s student) says(){
//	fmt.Println("I am just a regular student")
//}
//
//func (cr classrep) says(){
//	fmt.Println("I am a Class representive")
//}
//
//func main(){
//	p1:=student{
//		first:"Raman",
//		last:"Singh",
//		rollno:2,
//	}
//	p2:=classrep{
//		student:student{
//			first:"aman",
//			last:"sharma",
//		},
//		classrep:true,
//	}
//	p1.says()
//	p2.says()
//}
//
//struct pointer
//
//import "fmt"
//
//type person struct{
//	name string
//	age int
//}
//
//func main(){
//	p1:= &person{"ravi",24}
//	fmt.Println(p1)
//	fmt.Printf("%T",p1)//type
//	fmt.Println(p1.age)
//	fmt.Println(p1.name)
//}
//
//Json EXported /marshal
//
//import "fmt"
//import "encoding/json"
//type person struct{
//	Fname string //uppercase
//	Lname string //uppercase
//	Age int //uppercase
//	notExported string
//}
//
//func main(){
//	p1:=person{"Rama","Tiwari",25,"007"}
//	bs,_:=json.Marshal(p1)
//	fmt.Println(bs)
////	fmt.Println(v)shows nil
//	fmt.Println(string(bs))
//}
//
//Json tags
//import (
//	"fmt"
//	"encoding/json"
//)
//
//type person struct{
//	Fname string
//	Lname string `json:"-"`//json tag to not include , also notice we are using `` here and not ''
//	Age int `json:"wisdom score"`
//}
//
//func main(){
//	p1:=person{"Arun","Verma",42}
//	bs,_:=json.Marshal(p1)
//	fmt.Println(bs)
//	fmt.Println(string(bs))
//
//}
//
//Unmarshal
//import (
//	"fmt"
//	"encoding/json"
//)
//type person struct{
//	First string
//	Last string
//	Age int
//}
//
//func main(){
//	var p1 person
//	bs:=[]byte(`{"First":"Arun","Last":"Jaithley","Age":44}`)
//	json.Unmarshal(bs,&p1)
//	fmt.Println(p1.First,p1.Last)
//}
//
//Encode http://mindbowser.com/golang-data-structures-2/
//import (
//	"os"
//	"encoding/json"
//)
//type person struct{
//	First string
//	Last string
//	Age int
//}
//
//func main(){
//	p1:=person{"James","Bond",27}
//	json.NewEncoder(os.Stdout).Encode(p1)
//}
//
//Decode
//import (
//	"fmt"
//	"strings"
//	"encoding/json"
//)
//
//type person struct{
//	First string
//	Last string
//	Age int
//	notExported int
//}
//
//func main(){
//	var p1 person
//	rdr:=strings.NewReader(`{"First":"James"}`)
//	json.NewDecoder(rdr).Decode(&p1)
//	fmt.Println(p1.First)
//
//}
//
//import "fmt"
//
//type Square struct{ 						//define type of structure
//	side float64
//}
//
//func (z Square) area() float64{		//make method for  above type structure
//	return z.side * z.side
//}
//
//type Shape interface{					//make type interface
//	area() float64				//ye method hona chaiye un function mai jaha apne ko implementation chaiye is interface ki
//}
//
//func info(z Shape){						//Use above type interface to create a function so as to implement it in different form  in main func
//	fmt.Println(z)
//	fmt.Println(z.area())				//write the method created before
//}
//
//func main(){
//	s:=Square{10}
//	info(s)
//}
//
//import (
//	"fmt"
//	"math"
//)
//
//type Square struct{
//	side float64
//}
//type Circle struct{
//	radius float64
//}
//
//func (s Square) area() float64{
//	return s.side * s.side
//}
//
//func (c Circle) area() float64{
//	return 2*math.Pi *c.radius *c.radius
//}
//
//type Shape interface{
//	area() float64
//}
//
//func info (z Shape){
//	fmt.Println(z.area())
//}
//func main(){
//	s:=Square{10}
//	c:=Circle{5}
//	info(s)
//	info(c)
//}
//
//1.define types of elements
//2.make methods of above with same signature for each
//3.create a single interface type with the common signature from above in it
//4.make function with passing the interface in its parameter
//5.call in main
//
//import (
//"fmt"
//"sort"
//)
//
//type people []string
//
//func(p people) Len() int{ return len(p)}
//func(p people) Swap(i,j int){ p[i],p[j]=p[j],p[i]}
//func(p people) Less(i,j int)bool{return p[i]<p[j]}
//
//func main(){
//	studygroup:=people{"ram","ravi","sham","abhi"}
//	sort.Sort(studygroup)
//	fmt.Println(studygroup)
//}
//
//Sort a Variable not a type
//import(
//	"fmt"
//	"sort"
//)
//
//func main(){
//	s:=[]string{"rahul","ravi","raman","aman","gagan"}
//	sort.StringSlice(s).Sort()
//	fmt.Println(s)
//}
//
//Sort in reverse
//import (
//	"fmt"
//	"sort"
//)
//
//func main(){
//	t:=[]string{"rahul","abhi","ravi","sham"}
//	sort.Sort(sort.StringSlice(t))
//	fmt.Println(t)
//	sort.Sort(sort.Reverse(sort.StringSlice(t)))//make it other way around decending
//	fmt.Println(t)
//}
//
//Sor slice of ints
//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	n := []int{42, 56, 8, 45, 67, 23, 54}
//	sort.Ints(n)
//	fmt.Println(n)
//	//	sort.IntSlice(n).Sort()
//	//	fmt.Println(n)
//	//	sort.Sort(sort.IntSlice(n))
//	//	fmt.Println(n)
//	sort.Sort(sort.Reverse(sort.IntSlice(n)))
//	fmt.Println("reverse", n)
//
//}
//EMPTY INTERFACES //Type interface
//import "fmt"
//
//type vehical interface{}
//
//type car struct {
//	vehical
//	wheels int
//}
//type jet struct {
//	vehical
//	jets bool
//}
//type boat struct {
//	vehical
//	length int
//}
//
//func main() {
//	a := car{}
//	b := jet{}
//	c := boat{}
//
//	d := []vehical{a, b, c}
//
//	fmt.Println(d)
//}
//
//Method Sets
//import(
//	"fmt"
//	"math"
//)
//
//type Circle struct{
//	radius float64
//}
//
//type Shape interface{
//	area() float64
//}
//
//func (c Circle) area()float64{
//	return math.Pi *2 *c.radius *c.radius
//}
//
//func info(s Shape){
//	fmt.Println(s.area())
//}
//
//func main(){
//	d:=Circle{5}
//	info(&d)
//}
//Assertion
//import "fmt"
//
//func main(){
//	var name interface{}="Sydney"
//	str,ok:=name.(string) //assertion
//	if ok{
//		fmt.Printf("%q \n",str)
//	}else{
//		fmt.Printf("value is not a string\n")
//	}
//}
//
//Concurrency
//import (
//	"fmt"
//	"sync"
//	"time"
//	"runtime"
//)
//
//var wg sync.WaitGroup
//
//func init(){
//	runtime.GOMAXPROCS(runtime.NumCPU())
//}
//func main() {
//	wg.Add(2) //no of function
//	go foo()
//	//	foo()
//	//	bar()
//	go bar()
//	wg.Wait()
//}
//
//func foo() {
//	for i := 1; i < 50; i++ {
//		fmt.Println("foo", i)
//		time.Sleep(time.Duration(3 * time.Millisecond))
//	}
//	wg.Done()
//}
//
//func bar() {
//	for j := 1; j < 50; j++ {
//		fmt.Println("bar", j)
//		time.Sleep(time.Duration(3 * time.Millisecond))
//	}
//	wg.Done()
//}
//
//Race Condition
//import (
//	"fmt"
//	"time"
//	"sync"
//)
//
//var counter int
//var wg sync.WaitGroup
//
//func main(){
//	wg.Add(2)
//	go incrementor("Foo : ")
//	go incrementor("Bar : ")
//	wg.Wait()
//}
//
//func incrementor(s string){
//	for i:=0;i<20;i++{
//		x:=counter
//		x++
//		time.Sleep(time.Duration(3 * time.Millisecond))
//		counter=x
//		fmt.Print(s,i)
//		fmt.Println("  Counter ", counter)
//	}
//	wg.Done()
//}
//
//Mutex
//import (
// 	"fmt"
//	"time"
//	"sync"
//)
//
//var counter int
//var wg sync.WaitGroup
//var mutex sync.Mutex //for mutex
//
//func main(){
//	wg.Add(2)
//	go incrementor("Foo : ")
//	go incrementor("Bar : ")
//	wg.Wait()
//}
//
//func incrementor(s string){
//	for i:=0; i<20 ;i++{
//		time.Sleep(time.Duration(3 * time.Millisecond))
//		mutex.Lock()  //locking so that only one thread can use it at a time
//		counter++
//		fmt.Println(s,i," Counter ",counter)
//		mutex.Unlock()
//	}
//	wg.Done()
//}
//
//Atomicity
//
//import (
//	"fmt"
//	"time"
//	"sync/atomic"
//	"sync"
//	"math/rand"
//)
//var counter int64
//var wg sync.WaitGroup
//
//func main(){
//	wg.Add(2)
//	go incrementor(" Foo : ")
//	go incrementor(" Bar : ")
//	wg.Wait()
//}
//
//func incrementor(s string){
//	for i:=0;i<20; i++{
//		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
//		atomic.AddInt64(&counter,1)
//		fmt.Println(s,i," Counter: ",counter)
//	}
//	wg.Done()
//}
//import (
//    "fmt"
//    "math/rand"
//    "sync"
//    "sync/atomic"
//    "time"
//)
//
//var wg sync.WaitGroup
//var counter int64
//
//func main() {
//    wg.Add(2)
//    go incrementor("Foo:")
//    go incrementor("Bar:")
//    wg.Wait()
//    fmt.Println("Final Counter:", counter)
//}
//
//func incrementor(s string) {
//    for i := 0; i < 20; i++ {
//        time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
//        atomic.AddInt64(&counter, 1)
//        fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter)) // access without race
//    }
//    wg.Done()
//}
//
//Channels
//import (
//	"fmt"
//	"time"
//)
//
//func main(){
//	c:=make(chan int)
//
//	go func(){
//		for i:=1;i<10;i++{
//			fmt.Println("before value passed to channel : ",i) //the value comes out in two
//			c<-i
//		}
//	}()
//	go func(){
//		for {
//			fmt.Println(<-c)
//		}
//	}()
//	time.Sleep(time.Second)
//}
//
//Using of range cnd close
//
//import "fmt"
//
//func main(){
//	c:=make(chan int)
//
//	go func(){
//		for i:=1;i<10;i++{
//			c<-i
//		}
//		close(c) //closes after loop ends passing value to channel
//	}()
//	for n:=range c{
//		fmt.Println(n)
//	}
//}
//
//N to 1 (Multiple Goroutines to 1 channel)
//import "fmt"
//import "sync"
//
//var wg sync.WaitGroup
//
//func main(){
//	c:=make (chan int)
//	wg.Add(2)
//	go func(){
//		for i :=1;i<10;i++{
//			fmt.Println("before value passed to channel : ",i)
//			c<-i
//		}
//		wg.Done()
//	}()
//	go func(){
//		for i :=1;i<10;i++{
//			fmt.Println("value passed to channel : ",i)
//			c<-i
//		}
//		wg.Done()
//	}()
//
//	go func(){
//		wg.Wait()
//		close(c)
//	}()
//
//	for n:=range c{
//		fmt.Println(n)
//	}
//}
//Semaphore:Making a value (Here Channel values) to be used as a say flag
//import "fmt"
//
//func main(){
//	c:=make(chan int)
//	done:=make(chan bool)
//
//	go func(){
//		for i:=1;i<10;i++{
//			c<-i
//		}
//		done<-true
//	}()
//	go func (){
//		for i:=1;i<10;i++{
//			c<-i
//		}
//		done <-true
//	}()
//
//	//wrong way to do it
//	// <-done
//	// <-done
//	// close(c)
//
//	go func(){
//		<-done //throwing off
//		<-done //throwing off
//		close(c)
//	}()
//	for n:= range c{
//		fmt.Println(n)
//	}
//}
//
//N-to-1 but with many goroutines(basically more than 3)
//import "fmt"
//
//func main(){
//	n:=10
//	c:=make(chan int)
//	done:=make(chan bool)
//
//	for i:=1;i<=n;i++{ //repeat go func
//		go func(){
//			for i:=1;i<10;i++{
//				c<-i
//			}
//			done<-true
//		}()
//	}
//
//	go func (){
//		for i:=1;i<=n;i++{  //repeat same no of times "done" too
//			<-done
//		}
//		close(c)
//	}()
//
//	for n:=range c{
//		fmt.Println(n)
//	}
//}
//
//1-to-n
//import "fmt"
//
//func main(){
//	c:=make(chan int)
//	done:=make(chan bool)
//	go func (){
//		for i:=0;i<10;i++{
//			c<-i
//		}
//		close(c)
//	}()
//	go func(){
//		for n:=range c{
//			fmt.Println(n)
//		}
//		done<-true
//	}()
//	go func(){
//		for n:=range c{
//			fmt.Println(n)
//		}
//		done<-true
//	}()
//	<-done
//	<-done
//}
//
//Channels as arguments and returns
//import "fmt"
//
//func main() {
//	c := incrementor()
//	cSum := puller(c)
//	for n := range cSum {
//		fmt.Println(n)
//	}
//}
//
//func incrementor() chan int {
//	out := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			out <- i
//		}
//		close(out)
//	}()
//	return (out)
//}
//
//func puller(c chan int) chan int {
//	out2 := make(chan int)
//	go func() {
//		var Sum int
//		for n := range c {
//			Sum += n
//		}
//		out2 <- Sum
//		close(out2)
//	}()
//	return (out2)
//}
//
//Channels direction
//import "fmt"
//
//func main() {
//	c := incrementor()
//	for n := range  puller(c) {
//		fmt.Println(n)
//	}
//}
//
//func incrementor() <-chan int {
//	out := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			out <- i
//		}
//		close(out)
//	}()
//	return (out)
//}
//
//func puller(c <-chan int) <-chan int {
//	out2 := make(chan int)
//	go func() {
//		var Sum int
//		for n := range c {
//			Sum += n
//		}
//		out2 <- Sum
//		close(out2)
//	}()
//	return (out2)
//}
//
//Race Vs Deadlock/ Incrementor With Channels
//import "fmt"
//
//func main(){
//	c1:=incrementor("foo ")
//	c2:=incrementor("bar ")
//	c3:=puller(c1)
//	c4:=puller(c2)
//	fmt.Println("fnal Counter ",<-c4+<-c3)   //Query
//}
//func incrementor(s string)chan int{
//	out:=make(chan int)
//	go func (){
//		for i:=1;i<10;i++{
//			out<-i
//			fmt.Println(s,i)
//		}
//		close(out)
//	}()
//	return out
//}
//func puller(c chan int)chan int{
//	out2:=make(chan int)
//	go func(){
//		var sum int
//		for n:=range c{
//			sum = sum +n
//		}
//		out2<-sum
//		close(out2)
//	}()
//	return out2
//}
//
//Deadlock
//import "fmt"
//
//func main(){
//	c:=make(chan int)
//	c<-1
//	fmt.Println(<-c)
//}
//Deadlock removed
//import "fmt"
//
//func main(){
//	c:=make(chan int)
//	go func (){
//		c<-1
//	}()
//	fmt.Println(<-c)
//}
//
//Why Use Range
//
//import "fmt"
//
//func main(){
//	c:=make(chan int)
//	go func (){
//		for i:=0;i<10;i++{
//			c<-i
//		}
//		close(c)
//	}()
//	
////	fmt.Println(<-c)
//
////	for{
////		fmt.Println(<-c)
////	}
//	
//	for n:=range c{
//		fmt.Println(n)
//	}
//}
//
//Factorials via channel
//import "fmt"
//
//func main(){
//	f:=factorial(10)
//	for n:= range f{		
//		fmt.Println(n)
//	}
//}
//
//func factorial(n int)chan int{
//	out:=make(chan int)
//	go func(){
//		total:=1
//		for i:=n;i>0;i--{
//			total= total *i
//		}
//		out<-total
//		close(out)
//	}()
//	return out
//}
//
////PipeLines
//import "fmt"
//
//func main(){
////	c:=gen(2,3)
////	out:=sq(c)
//	//third
//// fmt.Println(<-out)
////	fmt.Println(<-out)
//	
////	for r:=range out{
////		fmt.Println(r)
////	}
//	
//	for r:=range sq(gen(2,3)){
//		fmt.Println(r)
//	}
//}
////first
//func gen(num ...int)chan int{
//	out:=make(chan int)
//	go func(){
//		for _,m:=range num{
//			out<-m
//		}
//		close(out)
//	}()
//	return out
//}
////second
//func sq(in chan int)chan int{
//	out:=make(chan int)
//	go func(){
//		for n:=range in{
//			out<-n*n
//		}
//		close(out)
//	}()
//	return out
//}
//
//Factorials Via Channel
//
//import "fmt"
//
//func main(){
//	in:= gen()
//	f:= factorial(in)
//	for n:= range f{
//		fmt.Println(n)
//	}
//}
//
//func gen() chan int{
//	out:=make(chan int)
//	go func (){
//		for i:=0;i<10;i++{
//			for j:=3;j<13;j++{
//				out<-j
//			}
//		}
//		close(out)
//	}()
//	return out
//}
//
//func factorial(in chan int) chan int{
//	out:=make(chan int)
//	go func(){
//		for n:=range in{
//			out<-fact(n)
//		}
//		close(out)
//	}()
//	return out
//}
//
//func fact(n int) int{
//	total:=1
//	for i:=n;i>0;i--{
//		total=total*i
//	}
//	return total
//}
//Fan In
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//func main(){
//	c:=FanIn(boring("Ann"),boring("John"))
//	for i:=0;i<10;i++{
//		fmt.Println(<-c)
//	}
//	fmt.Println("finished")
//}
//func boring(msg string) <-chan string{
//	c:=make(chan string)
//	go func(){
//		for i:=0;;i++{
//			c<-fmt.Sprintf("%s %d",msg,i)
//			time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
//		}
//	}()
//	return c
//}
//
////fan In
//func FanIn(input1,input2 <-chan string)<-chan string{
//	c:=make(chan string)
//	go func(){
//		for {
//			c<-<-input1
//		}
//	}()
//	go func(){
//		for {
//			c<-<-input2
//		}
//	}()
//	return c
//}
//
//Fan In
import "fmt"

func main(){
	in:=gen(2,3)
}

func gen(num ...int)chan int{
	fmt.Printf("%T",num)
}