# go-lessons
Trying out Golang

## Notes
Book being referred: An Introduction to Programming in Go by Caleb Doxsey

### Introduction 
import is similar to #include

main() is like the all universal main()

Exported functions have the first letter Capitalized (Println)
fmt.Println() or fmt.Printf() can be used for prints for stdout
fmt.Scanf() (same/similar as c) for reading input from stdin
### Types

#### Integers
uint8,uint16,uint32,uint64,int8,int16,int32,int64
int means signed integer
byte is same as uint8 (byte)
rune is same as int32 (normal int?)

uint,int and uintptr are arch specific types.

#### Floating Point
are inexact so don't use them for comparisions (same in other langs)

32 and 64-bit varients
In addition to numbers, has "Nan" to represent 0/0 and infinity.

float32,float64 are the datatypes. float64 is like double in Java
complex64 and complex128 for representing complex numbers. I should use this sometime.


#### Character
a simple char in a string "Hello World" is represented as a byte i.e. ASCII

#### String
len() - length of string
zero indexed
concat  can be done using + sign.
== can be used to compare two strings(values)

#### Booleans
&& and
!! or
! not
"bool" is a type which can hold the value "false" or "true"

### Variables
```go
var <name> <type> is the way to declare a variable
var <name> := <value> then type can be skipped; compiler will infer
<name> := <value>  var is also options (: denotes initialization)
```
You can declare multiple variables using following syntax; more readable in every function; multiple datatypes are allowed
```go
var (
	a = 42
	b = 3.14
)
var (
	a string //can also specify data-types
	b bool 
)
```
### Constants
```go
const <name>  <type> = <value> 
```
cannot be changed once declared
Above multiple variable declaration can be applied to constants too

### Arithematic
Normal operations + - * / %
Shorthand operators available += and others
++ -- standard increment and decrement operators

### Scoping
Global definitions accessible to all functions in the file
Normal function closures exist
Lexically scoped using blocks.

### Control Structures

#### if else
```go
if <condition> { //no () needed around the condition
} else if <condition> {
} else {
}
```


#### Loop
only one loop - for

eg- 
```go
i := 1
for i <= 10 { //only validation here; this is the while loop in go
	i = i+1
}
```
OR
```go
for i:= 1 ; i <=10 ; i++ { //normal standard for
}
```
#### switch
kinda cool because you don't need to write 'break';less debugging effort
```go
switch i {
case 42:
case 43:
default:
}
```

### Arrays

```go
var <name> [<size>]<type>
var x [4]int
x := [3]int{ 1,2,3}

```
Arrays are zero-indexed
len(array) will give you the length of the array
`Arrays are pass by value`
#### Iterating over Arrays
A special for loop exists for iterating over arrays
```go
for i,value := range <array> {
	//i is the current position
	//value is array[i]	
}
```
go compiler will not allow you to declare and not user these variables;
you can skip them by using the follow form
```go
for _,value := range <array> {
}
for i,_ := range <array> {
}
```
### Slices
Like arrays but allows variable length
Slices are backed by array of a fixed length.
The length of the slice can grow upto the size(called capacity) of the backing array
Most go libraries use slice rather than array
slice is pass by reference
```go 
var x []float64 //reference to a slice;needs to be instantiated/initialized
x:= make([]float64, 5 ,10)
len(x) //gives 5
cap(x) //gives 10

//more ways to intialize a size from an existing slice/array
slice := []int{11,22,32,42,546,2}
x := [low:high] //both can be skipped; low is start index(inclusive); high is last index(exclusive)
x := slice[:] //x has the entire array; same as [0:],[:],[0:len(slice)]
y := slice[0:2] //only first two elements
z := slice[ 2:] // array of all elements from index 2 till end
```
cap(array) will always be == len(array)
Refer https://blog.golang.org/slices for a better understanding

#### Slice in-built functions
two inbuilt functions - append and copy

```go 
//appendd
slice1 := []int{1,2,3}
slice2 := append(slice1,4,5) //kind of var args in Java
//notice how append returns a new slice; the backing array is different now
fmt.Println(slice1,slice2) //1,2,3,4,5
//copy
slice1 := []int{1,2,3}
slice2 := make([]int,2)
copy(slice2,slice1) //only first two elements will be copied
//copy returns the number of elements copied
```
TODO read more about slices

### Maps
Like dictionaries; 
Slices and maps are quite magical; not sure if user defined can be developed with similar style
```go
var x map[string]int //just a reference;needs instantiation/initialization
x := make(map[string]int)
x["hello"] = 42
len(x) // gives the number of mapings present in the map
//retrieve
i , ok := x["hello"] // both can be don't cares (_) //ok is boolean 

//init multiple:
test := map[string]int{
	"a" : 121,
	"b" : 23, //notice the ,
}
map can point to a map of some other values too
```

#### Slice in-built functions

```go
delete(x,"hello") //doesn't return anything;will do nothing if key no present
```
### Functions
Oh yeah!
Function call stack as in other languages
functions can return multiple values
Variadic functions :variable number of arguement)(varargs in Java)
Named return values are possible too
function can also return a function
function parameter area always pass by value;even arrays
```go
func <func-name>(<arg0> <type0>,<arg1> <type1> (<return-type1>,<return-type2>) {
}
//var args
func <func-name>(arg0 ...int) int{ } //can be invoked using a slice too

func test() (id int) {
	id = 42
	return //named return varaible; no need to explicitly specify vairalbe with return
}
```
#### Closures
can declare local functions inside functions which will have access to the local variables
also possible to assign a function to a variable (some thing like a function pointer)

### Defer,Panic & Recover
Panic and recover are go's exception handling
Defer seems something cool. Can it be dangerous?

#### Defer
Schedule a function call to run after the current function completes
used to free up resource (files)
keeps close near the open ; makes code less error prone
multiple return points in a function; then works fine too
```go 
func foo(){
	fmt.Println(1)
}
func bar(){
	fmt.Println(2
}
func main(){
	defere bar()
	foo()
}
//outputs 1 and then 2

//usecase-
f,_ := os.Open(filename)
defer f.close()
```
#### Panic & Recover
call panic() to throw an error
recover() will recover from the panic
TODO more details here

### Pointer
same like c
* and & operators
* is used to declare a pointer and also to dereference
& refers to the address of the following variable
```go
x := 4
var ptr *int
ptr = &x
```

#### new
Another way to  get a pointer
inbuilt function
takes a type as an argument and allocates enough memory
```go
ptr := new(int)
//ptr is of type *int
```
No Need to free up memory allocated by new.
Go is garbage collected

### Structs
similar like c
```go
//define
type Student struct {
	id int64
	name string
	address string
	//or
	name,address string
}
//initialize
var c Circle //like a pointer
//create
c = new(Student)
c := new(Student) //all in one
//to give values,like a constructor??
c := Student{ id: 42 , name:"John Doe",address: "Mars"}
c := Student{  42 ,"John Doe","Mars"}
c.id //gives 42
```
### Methods
Objected oriented way 
assigning methods to structs
put the type name between the keyword func and func-name
```go 
func (s *Student) getName() string {
	return s.name
}
```
Now you can call c.getName on the above student.

### Embedded types
Go's way of inheritance; is-a relation
also known as a embedded type
```go
type Student struct {
	Person //anonymous field
	id int64
}
//suppose Person has a method getName(),then
s := new(Student)
s.Person.getName()
//or
s.getName() //both are valid
```
### Interfaces
Like Java's interface
you don't have to define a struct as implements;
just implementing the methods defined in an interaface for a struct will be sufficient
compiler somehow figures out the relation
interfaces can also be used as variables as well as fields in other structs
```go
type Shape interface {
	area() float64
} 
//if circle and rectangle define this method then
var shape Shape
shape = Rectangle{3,4}
shape.area() //give 12
```
### Concurrency
The coolest feature probably. Caveats?
invoke a function like following & it will be invoked in a goroutine-a lightweight thread
```go
func f() {
}

go f() //invocation
foo() //this will be executed once the go routine is started; f() will be put on a seperate thread to be executed
```
main() itself is an implicit goroutine

### Channels
Inter goroutine communication and synchronization
if multiple goroutines are reading/writing from/to a channel then those read/writes will be synchronized by go
```go
var c  chan string - make(chan string) //a channel of string
//this has to be passed to goroutines which will read/write from it.
c <- "" //send to thechannel
msg := <-c //receive from channel ; blocking call
channel2  := make(chan string,10) //buffered async channel
```
